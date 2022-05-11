package keeper

import (
	"encoding/json"

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
func (k Keeper) ChainContracts(ctx sdk.Context) (res string) {
	k.paramstore.Get(ctx, types.KeyChainContracts, &res)
	return
}

func (k Keeper) ChainContractsParsed(ctx sdk.Context) types.ChainContractsParsed {
	var chainContractsStr string
	var chainContracts types.ChainContractsParsed
	k.paramstore.Get(ctx, types.KeyChainContracts, &chainContractsStr)
	err := json.Unmarshal([]byte(chainContractsStr), &chainContracts)

	if err != nil {
		panic(err)
	}

	return chainContracts
}
