package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/soupy-finance/noodle/x/bridge/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.ChainContracts(ctx),
	)
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

// ChainContracts returns the ChainContracts param
func (k Keeper) ChainContracts(ctx sdk.Context) (res map[string]string) {
	k.paramstore.Get(ctx, types.KeyChainContracts, &res)
	return
}
