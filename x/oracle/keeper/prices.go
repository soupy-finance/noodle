package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	staking "github.com/cosmos/cosmos-sdk/x/staking/types"
)

func (k Keeper) AggAssetPrice(ctx sdk.Context, asset string) uint64 {
	var price uint64
	var prices []uint64
	k.stakingKeeper.IterateBondedValidators(ctx, GetValidatorPrice(asset, &prices))

	// Process prices slice
	//

	return price
}

func GetValidatorPrice(asset string, prices *[]uint64) (fn func(index int64, validator staking.ValidatorI) (stop bool)) {
	// Get store
	//

	return func(index int64, validator staking.ValidatorI) (stop bool) {
		// Get price from store and update prices slice
		//

		return false
	}
}
