package keeper

import (
	"encoding/json"

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
func (k Keeper) Assets(ctx sdk.Context) (res string) {
	k.paramstore.Get(ctx, types.KeyAssets, &res)
	return
}

func (k Keeper) AssetsParsed(ctx sdk.Context) types.AssetsParsed {
	var assetsStr string
	var assets types.AssetsParsed
	k.paramstore.Get(ctx, types.KeyAssets, &assetsStr)
	err := json.Unmarshal([]byte(assetsStr), &assets)

	if err != nil {
		panic(err)
	}

	return assets
}
