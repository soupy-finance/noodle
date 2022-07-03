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
	addr1 := sample.AccAddress()
	addr2 := sample.AccAddress()
	balances := map[string]map[string]sdk.Int{
		addr1: {
			"eth":  sdk.ZeroInt(),
			"usdc": sdk.ZeroInt(),
		},
		addr2: {
			"eth":  sdk.ZeroInt(),
			"usdc": sdk.ZeroInt(),
		},
	}

	tests := []struct {
		name          string
		msg           types.MsgCreateOrder
		msgs          []types.MsgCreateOrder
		err           error
		init          func()
		check         func(*testing.T, keeper.Keeper, types.MsgServer, context.Context, *types.MsgCreateOrder, *types.MsgCreateOrderResponse)
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
			check: func(t *testing.T, k keeper.Keeper, msgServer types.MsgServer, goCtx context.Context, msg *types.MsgCreateOrder, res *types.MsgCreateOrderResponse) {
				ctx := sdk.UnwrapSDKContext(goCtx)
				side, _ := types.NewSide('b')
				book, _ := k.GetVirtualBook(ctx, "eth-usdc", side)

				require.Equal(t, 1, len(book.Levels))
				require.Equal(t, 1, len(book.Levels[0].Orders))
				require.Equal(t, "2000.000000000000000000", book.Levels[0].Orders[0].Price.String())
				require.Equal(t, "1.000000000000000000", book.Levels[0].Orders[0].Quantity.String())
				require.Equal(t, "2000.000000000000000000", book.BestPrice().String())
			},
			bankKeeper: BankKeeper{
				_getBalance: func(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin {
					var quantityDec sdk.Dec

					switch denom {
					case "usdc":
						quantityDec = sdk.MustNewDecFromStr("2000")
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
		{
			name: "limit match 1 eth at 2000",
			msgs: []types.MsgCreateOrder{
				{
					Creator:   addr1,
					Market:    "eth-usdc",
					Side:      false,
					OrderType: "limit",
					Price:     "2000",
					Quantity:  "1",
				},
				{
					Creator:   addr2,
					Market:    "eth-usdc",
					Side:      true,
					OrderType: "limit",
					Price:     "2000",
					Quantity:  "1",
				},
			},
			init: func() {
				balances[addr1]["usdc"] = sdk.NewIntFromBigInt(sdk.MustNewDecFromStr("2000").BigInt())
				balances[addr2]["eth"] = sdk.NewIntFromBigInt(sdk.MustNewDecFromStr("1").BigInt())
			},
			check: func(t *testing.T, k keeper.Keeper, msgServer types.MsgServer, goCtx context.Context, msg *types.MsgCreateOrder, res *types.MsgCreateOrderResponse) {
				ctx := sdk.UnwrapSDKContext(goCtx)
				side, _ := types.NewSide('b')
				book, _ := k.GetVirtualBook(ctx, "eth-usdc", side)
				addr1Eth := sdk.NewDecFromIntWithPrec(balances[addr1]["eth"], sdk.Precision).String()
				addr2Usdc := sdk.NewDecFromIntWithPrec(balances[addr2]["usdc"], sdk.Precision).String()
				addr1Usdc := sdk.NewDecFromIntWithPrec(balances[addr1]["usdc"], sdk.Precision).String()
				addr2Eth := sdk.NewDecFromIntWithPrec(balances[addr2]["eth"], sdk.Precision).String()

				require.Equal(t, 0, len(book.Levels))
				require.Equal(t, sdk.MustNewDecFromStr("1").String(), addr1Eth)
				require.Equal(t, sdk.MustNewDecFromStr("2000").String(), addr2Usdc)
				require.Equal(t, sdk.ZeroDec().String(), addr1Usdc)
				require.Equal(t, sdk.ZeroDec().String(), addr2Eth)
			},
			bankKeeper: BankKeeper{
				_getBalance: func(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin {
					var quantityDec sdk.Dec

					switch denom {
					case "usdc":
						quantityDec = sdk.MustNewDecFromStr("2000")
					case "eth":
						quantityDec = sdk.MustNewDecFromStr("1")
					default:
						quantityDec = sdk.ZeroDec()
					}

					return sdk.NewCoin(denom, sdk.NewIntFromBigInt(quantityDec.BigInt()))
				},
				_sendCoins: func(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amts sdk.Coins) error {
					for i := range amts {
						coin := amts[i]
						fromAddrBech32 := fromAddr.String()
						toAddrBech32 := toAddr.String()
						balances[fromAddrBech32][coin.Denom] = balances[fromAddrBech32][coin.Denom].Sub(coin.Amount)
						balances[toAddrBech32][coin.Denom] = balances[toAddrBech32][coin.Denom].Add(coin.Amount)
					}

					return nil
				},
				_sendCoinsFromModuleToAccount: func(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amts sdk.Coins) error {
					for i := range amts {
						coin := amts[i]
						balances[recipientAddr.String()][coin.Denom] = balances[recipientAddr.String()][coin.Denom].Add(coin.Amount)
					}

					return nil
				},
				_sendCoinsFromAccountToModule: func(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amts sdk.Coins) error {
					for i := range amts {
						coin := amts[i]
						amtStr := coin.Amount.String()
						balStr := balances[senderAddr.String()][coin.Denom].String()
						_ = amtStr
						_ = balStr
						balances[senderAddr.String()][coin.Denom] = balances[senderAddr.String()][coin.Denom].Sub(coin.Amount)
					}

					return nil
				},
			},
		},
		{
			name: "limit don't match 1 eth at 2000",
			msgs: []types.MsgCreateOrder{
				{
					Creator:   sample.AccAddress(),
					Market:    "eth-usdc",
					Side:      false,
					OrderType: "limit",
					Price:     "2000",
					Quantity:  "1",
				},
				{
					Creator:   sample.AccAddress(),
					Market:    "eth-usdc",
					Side:      true,
					OrderType: "limit",
					Price:     "2001",
					Quantity:  "1",
				},
			},
			check: func(t *testing.T, k keeper.Keeper, msgServer types.MsgServer, goCtx context.Context, msg *types.MsgCreateOrder, res *types.MsgCreateOrderResponse) {
				ctx := sdk.UnwrapSDKContext(goCtx)
				bidSide, _ := types.NewSide('b')
				askSide, _ := types.NewSide('a')
				bidBook, _ := k.GetVirtualBook(ctx, "eth-usdc", bidSide)
				askBook, _ := k.GetVirtualBook(ctx, "eth-usdc", askSide)

				require.Equal(t, 1, len(bidBook.Levels))
				require.Equal(t, 1, len(bidBook.Levels[0].Orders))
				require.Equal(t, "2000.000000000000000000", bidBook.Levels[0].Orders[0].Price.String())
				require.Equal(t, "1.000000000000000000", bidBook.Levels[0].Orders[0].Quantity.String())
				require.Equal(t, "2000.000000000000000000", bidBook.BestPrice().String())

				require.Equal(t, 1, len(askBook.Levels))
				require.Equal(t, 1, len(askBook.Levels[0].Orders))
				require.Equal(t, "2001.000000000000000000", askBook.Levels[0].Orders[0].Price.String())
				require.Equal(t, "1.000000000000000000", askBook.Levels[0].Orders[0].Quantity.String())
				require.Equal(t, "2001.000000000000000000", askBook.BestPrice().String())
			},
			bankKeeper: BankKeeper{
				_getBalance: func(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin {
					var quantityDec sdk.Dec

					switch denom {
					case "usdc":
						quantityDec = sdk.MustNewDecFromStr("2000")
					case "eth":
						quantityDec = sdk.MustNewDecFromStr("1")
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
			k, msgServer, ctx := setupMsgServer(t, tt.bankKeeper, tt.stakingKeeper)
			var res *types.MsgCreateOrderResponse
			var err error

			if tt.init != nil {
				tt.init()
			}

			if len(tt.msgs) == 0 {
				res, err = msgServer.CreateOrder(ctx, &tt.msg)
			} else {
				for _, msg := range tt.msgs {
					res, err = msgServer.CreateOrder(ctx, &msg)
				}
			}

			_ = res
			require.Equal(t, tt.err, err)

			if tt.check != nil {
				if len(tt.msgs) == 0 {
					tt.check(t, k, msgServer, ctx, &tt.msg, res)
				} else {
					tt.check(t, k, msgServer, ctx, &tt.msgs[0], res)
				}
			}
		})
	}
}

func TestMsgCreateOrderMsg_MarketBuy(t *testing.T) {
	tests := []struct {
		name          string
		msg           types.MsgCreateOrder
		msgs          []types.MsgCreateOrder
		err           error
		check         func(*testing.T, keeper.Keeper, types.MsgServer, context.Context, *types.MsgCreateOrder, *types.MsgCreateOrderResponse)
		bankKeeper    BankKeeper
		stakingKeeper StakingKeeper
	}{
		{
			name: "market buy 1 eth at 2000 insufficient liquidity",
			msg: types.MsgCreateOrder{
				Creator:   sample.AccAddress(),
				Market:    "eth-usdc",
				Side:      false,
				OrderType: "market",
				Price:     "2000",
				Quantity:  "1",
			},
			bankKeeper: BankKeeper{
				_getBalance: func(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin {
					return sdk.NewCoin(denom, sdk.ZeroInt())
				},
			},
			err: types.InsufficientLiquidity,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k, msgServer, ctx := setupMsgServer(t, tt.bankKeeper, tt.stakingKeeper)
			var res *types.MsgCreateOrderResponse
			var err error

			if len(tt.msgs) == 0 {
				res, err = msgServer.CreateOrder(ctx, &tt.msg)
			} else {
				for _, msg := range tt.msgs {
					res, err = msgServer.CreateOrder(ctx, &msg)
				}
			}

			_ = res
			require.Equal(t, tt.err, err)

			if tt.check != nil {
				if len(tt.msgs) == 0 {
					tt.check(t, k, msgServer, ctx, &tt.msg, res)
				} else {
					tt.check(t, k, msgServer, ctx, &tt.msgs[0], res)
				}
			}
		})
	}
}

type BankKeeper struct {
	_spendableCoins               func(sdk.Context, sdk.AccAddress) sdk.Coins
	_getBalance                   func(sdk.Context, sdk.AccAddress, string) sdk.Coin
	_sendCoins                    func(sdk.Context, sdk.AccAddress, sdk.AccAddress, sdk.Coins) error
	_sendCoinsFromModuleToAccount func(sdk.Context, string, sdk.AccAddress, sdk.Coins) error
	_sendCoinsFromAccountToModule func(sdk.Context, sdk.AccAddress, string, sdk.Coins) error
}

func (k BankKeeper) SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins {
	return k._spendableCoins(ctx, addr)
}
func (k BankKeeper) GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin {
	return k._getBalance(ctx, addr, denom)
}
func (k BankKeeper) SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amts sdk.Coins) error {
	return k._sendCoins(ctx, fromAddr, toAddr, amts)
}
func (k BankKeeper) SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error {
	return k._sendCoinsFromModuleToAccount(ctx, senderModule, recipientAddr, amt)
}
func (k BankKeeper) SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error {
	return k._sendCoinsFromAccountToModule(ctx, senderAddr, recipientModule, amt)
}

type StakingKeeper struct{}
