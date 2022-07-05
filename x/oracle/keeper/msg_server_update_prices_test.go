package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	staking "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/soupy-finance/noodle/testutil/sample"
	"github.com/soupy-finance/noodle/x/oracle/keeper"
	"github.com/soupy-finance/noodle/x/oracle/types"
	"github.com/stretchr/testify/require"
)

func TestMsgUpdatePrices(t *testing.T) {
	valAccAddress := sample.AccAddress()

	tests := []struct {
		name          string
		msg           types.MsgUpdatePrices
		msgs          []types.MsgUpdatePrices
		do            func(*testing.T, keeper.Keeper, types.MsgServer, context.Context)
		err           error
		check         func(*testing.T, keeper.Keeper, types.MsgServer, context.Context, *types.MsgUpdatePrices, *types.MsgUpdatePricesResponse)
		bankKeeper    BankKeeper
		stakingKeeper StakingKeeper
	}{
		{
			name: "update prices not a validator",
			msg: types.MsgUpdatePrices{
				Creator: valAccAddress,
				Data:    "{\"eth\": \"1\"}",
			},
			err: types.NotValidator,
			stakingKeeper: StakingKeeper{
				_getValidator: func(ctx sdk.Context, address sdk.ValAddress) (validator staking.Validator, found bool) {
					return validator, false
				},
				_iterateBondedValidatorsByPower: func(ctx sdk.Context, fn func(index int64, val staking.ValidatorI) bool) {
					account := sample.AccAddress()
					valAddr, _ := sdk.ValAddressFromBech32(account)
					fn(0, staking.Validator{OperatorAddress: sdk.MustBech32ifyAddressBytes("cosmosvaloper", valAddr)})
				},
				_powerReduction: func(ctx sdk.Context) sdk.Int {
					return sdk.NewIntFromUint64(1000000)
				},
			},
		},
		{
			name: "update prices success",
			msg: types.MsgUpdatePrices{
				Creator: valAccAddress,
				Data:    "{\"eth\": \"1\"}",
			},
			do: func(t *testing.T, k keeper.Keeper, msgServer types.MsgServer, goCtx context.Context) {
				ctx := sdk.UnwrapSDKContext(goCtx)
				k.UpdateAndSavePrices(ctx)
			},
			check: func(t *testing.T, k keeper.Keeper, msgServer types.MsgServer, goCtx context.Context, msg *types.MsgUpdatePrices, res *types.MsgUpdatePricesResponse) {
				ctx := sdk.UnwrapSDKContext(goCtx)
				prices := k.GetAssetPrices(ctx, []string{"eth"})

				require.Equal(t, "1.000000000000000000", prices[0])
			},
			stakingKeeper: StakingKeeper{
				_getValidator: func(ctx sdk.Context, address sdk.ValAddress) (validator staking.Validator, found bool) {
					return staking.Validator{
						Status: 3,
						Tokens: sdk.OneInt().Mul(sdk.DefaultPowerReduction),
					}, true
				},
				_iterateBondedValidatorsByPower: func(ctx sdk.Context, fn func(index int64, val staking.ValidatorI) bool) {
					accAddr, _ := sdk.AccAddressFromBech32(valAccAddress)
					valAddr := sdk.ValAddress(accAddr.Bytes())
					prefix := sdk.GetConfig().GetBech32ValidatorAddrPrefix()
					fn(0, staking.Validator{OperatorAddress: sdk.MustBech32ifyAddressBytes(prefix, valAddr)})
				},
				_powerReduction: func(ctx sdk.Context) sdk.Int {
					return sdk.NewIntFromUint64(1000000)
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k, msgServer, ctx := setupMsgServer(t, tt.bankKeeper, tt.stakingKeeper)
			var res *types.MsgUpdatePricesResponse
			var err error

			if len(tt.msgs) == 0 {
				res, err = msgServer.UpdatePrices(ctx, &tt.msg)
			} else {
				for _, msg := range tt.msgs {
					res, err = msgServer.UpdatePrices(ctx, &msg)
				}
			}

			if tt.do != nil {
				tt.do(t, k, msgServer, ctx)
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
