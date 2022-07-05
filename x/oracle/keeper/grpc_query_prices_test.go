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

func TestPricesQuery(t *testing.T) {
	tests := []struct {
		name          string
		req           types.QueryPricesRequest
		reqs          []types.QueryPricesRequest
		err           error
		check         func(*testing.T, keeper.Keeper, context.Context, *types.QueryPricesRequest, *types.QueryPricesResponse)
		bankKeeper    BankKeeper
		stakingKeeper StakingKeeper
	}{
		{
			name: "test prices query",
			req: types.QueryPricesRequest{
				Assets: []string{"eth"},
			},
			check: func(t *testing.T, k keeper.Keeper, ctx context.Context, req *types.QueryPricesRequest, res *types.QueryPricesResponse) {
				require.Equal(t, "[\"0.000000000000000000\"]", res.Data)
			},
			bankKeeper: BankKeeper{},
			stakingKeeper: StakingKeeper{
				_iterateBondedValidatorsByPower: func(ctx sdk.Context, fn func(index int64, val staking.ValidatorI) bool) {
					account := sample.AccAddress()
					valAddr, _ := sdk.ValAddressFromBech32(account)
					fn(0, staking.Validator{OperatorAddress: sdk.MustBech32ifyAddressBytes("cosmosvaloper", valAddr)})
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k, ctx := setupKeeper(t, tt.bankKeeper, tt.stakingKeeper)
			var res *types.QueryPricesResponse
			var err error

			if len(tt.reqs) == 0 {
				res, err = k.Prices(ctx, &tt.req)
			} else {
				for _, req := range tt.reqs {
					res, err = k.Prices(ctx, &req)
				}
			}

			_ = res
			require.Equal(t, tt.err, err)

			if tt.check != nil {
				if len(tt.reqs) == 0 {
					tt.check(t, k, ctx, &tt.req, res)
				} else {
					tt.check(t, k, ctx, &tt.reqs[0], res)
				}
			}
		})
	}
}

type BankKeeper struct {
	_spendableCoins func(sdk.Context, sdk.AccAddress) sdk.Coins
}

func (k BankKeeper) SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins {
	return k._spendableCoins(ctx, addr)
}

type StakingKeeper struct {
	_getValidator                   func(sdk.Context, sdk.ValAddress) (staking.Validator, bool)
	_iterateBondedValidatorsByPower func(sdk.Context, func(int64, staking.ValidatorI) bool)
	_powerReduction                 func(sdk.Context) sdk.Int
}

func (k StakingKeeper) GetValidator(ctx sdk.Context, address sdk.ValAddress) (validator staking.Validator, found bool) {
	return k._getValidator(ctx, address)
}
func (k StakingKeeper) IterateBondedValidatorsByPower(ctx sdk.Context, fn func(index int64, validator staking.ValidatorI) (stop bool)) {
	k._iterateBondedValidatorsByPower(ctx, fn)
}
func (k StakingKeeper) PowerReduction(ctx sdk.Context) sdk.Int {
	return k._powerReduction(ctx)
}
