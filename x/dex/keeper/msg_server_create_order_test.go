package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/soupy-finance/noodle/testutil/sample"
	"github.com/soupy-finance/noodle/x/dex/keeper"
	"github.com/soupy-finance/noodle/x/dex/types"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateOrderMsg_LimitBuy(t *testing.T) {
	tests := []struct {
		name          string
		msg           types.MsgCreateOrder
		err           error
		check         func(*testing.T, keeper.Keeper, types.MsgServer, context.Context)
		bankKeeper    BankKeeper
		stakingKeeper StakingKeeper
	}{
		{
			name: "limit buy 1 eth at 2000 insufficient balance",
			msg: types.MsgCreateOrder{
				Creator:   sample.AccAddress(),
				Market:    "eth-usdc",
				Side:      false,
				OrderType: "limit",
				Price:     "2000",
				Quantity:  "1",
			},
			bankKeeper: BankKeeper{
				_getBalance: func(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin {
					return sdk.NewCoin(denom, sdk.ZeroInt())
				},
			},
			err: types.InsufficientBalance,
		},
		{
			name: "limit buy 1 eth at 2000 sufficient balance",
			msg: types.MsgCreateOrder{
				Creator:   sample.AccAddress(),
				Market:    "eth-usdc",
				Side:      false,
				OrderType: "limit",
				Price:     "2000",
				Quantity:  "1",
			},
			check: func(t *testing.T, k keeper.Keeper, msgServer types.MsgServer, goCtx context.Context) {
				ctx := sdk.UnwrapSDKContext(goCtx)
				side, _ := types.NewSide('b')
				book, _ := k.GetVirtualBook(ctx, "eth-usdc", side)

				require.Equal(t, len(book.Levels), 1)
				require.Equal(t, len(book.Levels[0].Orders), 1)
				require.Equal(t, "2000.000000000000000000", book.Levels[0].Orders[0].Price.String())
				require.Equal(t, "1.000000000000000000", book.Levels[0].Orders[0].Quantity.String())
				require.Equal(t, book.BestPrice().String(), "2000.000000000000000000")
			},
			bankKeeper: BankKeeper{
				_getBalance: func(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin {
					var quantityDec sdk.Dec

					switch denom {
					case "usdc":
						quantityDec, _ = sdk.NewDecFromStr("2000")
					default:
						quantityDec = sdk.ZeroDec()
					}

					return sdk.NewCoin(denom, sdk.NewIntFromBigInt(quantityDec.BigInt()))
				},
				_sendCoins: func(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error {
					return nil
				},
				_sendCoinsFromModuleToAccount: func(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error {
					return nil
				},
				_sendCoinsFromAccountToModule: func(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error {
					return nil
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k, msgServer, ctx := setupMsgServerAndKeeper(t, tt.bankKeeper, tt.stakingKeeper)
			res, err := msgServer.CreateOrder(ctx, &tt.msg)

			_ = res
			require.Equal(t, tt.err, err)

			if tt.check != nil {
				tt.check(t, k, msgServer, ctx)
			}
		})
	}
}

type BankKeeper struct {
	_spendableCoins               func(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	_getBalance                   func(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
	_sendCoins                    func(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error
	_sendCoinsFromModuleToAccount func(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	_sendCoinsFromAccountToModule func(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
}

func (k BankKeeper) SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins {
	return k._spendableCoins(ctx, addr)
}
func (k BankKeeper) GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin {
	return k._getBalance(ctx, addr, denom)
}
func (k BankKeeper) SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error {
	return k._sendCoins(ctx, fromAddr, toAddr, amt)
}
func (k BankKeeper) SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error {
	return k._sendCoinsFromModuleToAccount(ctx, senderModule, recipientAddr, amt)
}
func (k BankKeeper) SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error {
	return k._sendCoinsFromAccountToModule(ctx, senderAddr, recipientModule, amt)
}

type StakingKeeper struct{}
