package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/soupy-finance/noodle/x/oracle/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.Assets(ctx),
	)
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

// Assets returns the Assets param
func (k Keeper) Assets(ctx sdk.Context) (res []string) {
	k.paramstore.Get(ctx, types.KeyAssets, &res)
	return
}
