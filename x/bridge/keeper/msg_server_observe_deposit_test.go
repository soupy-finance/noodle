package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	auth "github.com/cosmos/cosmos-sdk/x/auth/types"
	staking "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/soupy-finance/noodle/testutil/sample"
	"github.com/soupy-finance/noodle/x/bridge/keeper"
	"github.com/soupy-finance/noodle/x/bridge/types"
	"github.com/stretchr/testify/require"
)

func TestMsgObserveDeposit(t *testing.T) {
	depositorExtAddr := sample.AccAddress()
	observer1 := sample.AccAddress()
	observer2 := sample.AccAddress()

	tests := []struct {
		name          string
		msg           types.MsgObserveDeposit
		msgs          []types.MsgObserveDeposit
		err           error
		check         func(*testing.T, keeper.Keeper, types.MsgServer, context.Context, *types.MsgObserveDeposit, *types.MsgObserveDepositResponse)
		AccountKeeper AccountKeeper
		bankKeeper    BankKeeper
		stakingKeeper StakingKeeper
	}{
		{
			name: "observe deposit not validator",
			msg: types.MsgObserveDeposit{
				Creator:   observer1,
				ChainId:   "ethereum",
				Depositor: depositorExtAddr,
				DepositId: "1",
				Quantity:  "1",
				Asset:     "eth",
			},
			err: types.NotValidator,
			stakingKeeper: StakingKeeper{
				_getValidator: func(ctx sdk.Context, address sdk.ValAddress) (validator staking.Validator, found bool) {
					return validator, false
				},
			},
		},
		{
			name: "observe deposit invalid chain",
			msg: types.MsgObserveDeposit{
				Creator:   observer1,
				ChainId:   "fake chain",
				Depositor: depositorExtAddr,
				DepositId: "1",
				Quantity:  "1",
				Asset:     "eth",
			},
			err: types.InvalidChain,
			stakingKeeper: StakingKeeper{
				_getValidator: func(ctx sdk.Context, address sdk.ValAddress) (validator staking.Validator, found bool) {
					return staking.Validator{Status: 3}, true
				},
			},
		},
		{
			name: "observe deposit partial",
			msg: types.MsgObserveDeposit{
				Creator:   observer1,
				ChainId:   "ethereum",
				Depositor: depositorExtAddr,
				DepositId: "1",
				Quantity:  "1",
				Asset:     "eth",
			},
			check: func(t *testing.T, k keeper.Keeper, msgServer types.MsgServer, goCtx context.Context, msg *types.MsgObserveDeposit, res *types.MsgObserveDepositResponse) {
				ctx := sdk.UnwrapSDKContext(goCtx)
				depositor := k.ExternalAddressToDexAddress(ctx, msg.Depositor)
				depositKeyBytes := keeper.GetDepositKeyBytes(depositor, msg)
				observations, err := k.GetDepositObservations(ctx, depositKeyBytes, msg)

				require.Equal(t, nil, err)
				require.Equal(t, 1, len(observations))
			},
			stakingKeeper: StakingKeeper{
				_getValidator: func(ctx sdk.Context, address sdk.ValAddress) (validator staking.Validator, found bool) {
					return staking.Validator{
						Status: 3,
						Tokens: sdk.OneInt().Mul(sdk.DefaultPowerReduction),
					}, true
				},
				_getLastTotalPower: func(ctx sdk.Context) sdk.Int {
					return sdk.NewIntFromUint64(2)
				},
			},
		},
		{
			name: "observe deposit full",
			msgs: []types.MsgObserveDeposit{
				{
					Creator:   observer1,
					ChainId:   "ethereum",
					Depositor: depositorExtAddr,
					DepositId: "1",
					Quantity:  "1",
					Asset:     "eth",
				},
				{
					Creator:   observer2,
					ChainId:   "ethereum",
					Depositor: depositorExtAddr,
					DepositId: "1",
					Quantity:  "1",
					Asset:     "eth",
				},
			},
			check: func(t *testing.T, k keeper.Keeper, msgServer types.MsgServer, goCtx context.Context, msg *types.MsgObserveDeposit, res *types.MsgObserveDepositResponse) {
				ctx := sdk.UnwrapSDKContext(goCtx)
				depositor := k.ExternalAddressToDexAddress(ctx, msg.Depositor)
				depositKeyBytes := keeper.GetDepositKeyBytes(depositor, msg)
				observations, err := k.GetDepositObservations(ctx, depositKeyBytes, msg)

				require.Equal(t, nil, err)
				require.Equal(t, 0, len(observations))
			},
			stakingKeeper: StakingKeeper{
				_getValidator: func(ctx sdk.Context, address sdk.ValAddress) (validator staking.Validator, found bool) {
					return staking.Validator{
						Status: 3,
						Tokens: sdk.OneInt().Mul(sdk.DefaultPowerReduction),
					}, true
				},
				_getLastTotalPower: func(ctx sdk.Context) sdk.Int {
					return sdk.NewIntFromUint64(2)
				},
			},
			bankKeeper: BankKeeper{
				_mintCoins: func(ctx sdk.Context, asset string, amt sdk.Coins) error {
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
			name: "observe deposit no dup",
			msgs: []types.MsgObserveDeposit{
				{
					Creator:   observer1,
					ChainId:   "ethereum",
					Depositor: depositorExtAddr,
					DepositId: "1",
					Quantity:  "1",
					Asset:     "eth",
				},
				{
					Creator:   observer1,
					ChainId:   "ethereum",
					Depositor: depositorExtAddr,
					DepositId: "1",
					Quantity:  "1",
					Asset:     "eth",
				},
			},
			err: types.DupObservation,
			check: func(t *testing.T, k keeper.Keeper, msgServer types.MsgServer, goCtx context.Context, msg *types.MsgObserveDeposit, res *types.MsgObserveDepositResponse) {
				ctx := sdk.UnwrapSDKContext(goCtx)
				depositor := k.ExternalAddressToDexAddress(ctx, msg.Depositor)
				depositKeyBytes := keeper.GetDepositKeyBytes(depositor, msg)
				observations, err := k.GetDepositObservations(ctx, depositKeyBytes, msg)

				require.Equal(t, nil, err)
				require.Equal(t, 1, len(observations))
			},
			stakingKeeper: StakingKeeper{
				_getValidator: func(ctx sdk.Context, address sdk.ValAddress) (validator staking.Validator, found bool) {
					return staking.Validator{
						Status: 3,
						Tokens: sdk.OneInt().Mul(sdk.DefaultPowerReduction),
					}, true
				},
				_getLastTotalPower: func(ctx sdk.Context) sdk.Int {
					return sdk.NewIntFromUint64(2)
				},
			},
			bankKeeper: BankKeeper{
				_mintCoins: func(ctx sdk.Context, asset string, amt sdk.Coins) error {
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
			var res *types.MsgObserveDepositResponse
			var err error

			if len(tt.msgs) == 0 {
				res, err = msgServer.ObserveDeposit(ctx, &tt.msg)
			} else {
				for _, msg := range tt.msgs {
					res, err = msgServer.ObserveDeposit(ctx, &msg)
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

type AccountKeeper struct {
	_getAccount func(sdk.Context, sdk.AccAddress) auth.AccountI
}

func (k AccountKeeper) GetAccount(ctx sdk.Context, addr sdk.AccAddress) auth.AccountI {
	return k._getAccount(ctx, addr)
}

type BankKeeper struct {
	_spendableCoins               func(sdk.Context, sdk.AccAddress) sdk.Coins
	_mintCoins                    func(sdk.Context, string, sdk.Coins) error
	_burnCoins                    func(sdk.Context, string, sdk.Coins) error
	_sendCoinsFromModuleToAccount func(sdk.Context, string, sdk.AccAddress, sdk.Coins) error
	_sendCoinsFromAccountToModule func(sdk.Context, sdk.AccAddress, string, sdk.Coins) error
}

func (k BankKeeper) SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins {
	return k._spendableCoins(ctx, addr)
}
func (k BankKeeper) MintCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error {
	return k._mintCoins(ctx, moduleName, amt)
}
func (k BankKeeper) BurnCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error {
	return k._burnCoins(ctx, moduleName, amt)
}
func (k BankKeeper) SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error {
	return k._sendCoinsFromModuleToAccount(ctx, senderModule, recipientAddr, amt)
}
func (k BankKeeper) SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error {
	return k._sendCoinsFromAccountToModule(ctx, senderAddr, recipientModule, amt)
}

type StakingKeeper struct {
	_getValidator      func(sdk.Context, sdk.ValAddress) (staking.Validator, bool)
	_getLastTotalPower func(sdk.Context) sdk.Int
}

func (k StakingKeeper) GetValidator(ctx sdk.Context, address sdk.ValAddress) (validator staking.Validator, found bool) {
	return k._getValidator(ctx, address)
}
func (k StakingKeeper) GetLastTotalPower(ctx sdk.Context) sdk.Int {
	return k._getLastTotalPower(ctx)
}
