package keeper

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/soupy-finance/noodle/x/dex/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.Markets(ctx),
	)
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

// Markets returns the Markets param
func (k Keeper) Markets(ctx sdk.Context) (res string) {
	k.paramstore.Get(ctx, types.KeyMarkets, &res)
	return
}

func (k Keeper) MarketsParsed(ctx sdk.Context) types.MarketsParsed {
	var marketsStr string
	var markets types.MarketsParsed
	k.paramstore.Get(ctx, types.KeyMarkets, &marketsStr)
	err := json.Unmarshal([]byte(marketsStr), &markets)

	if err != nil {
		panic(err)
	}

	return markets
}
