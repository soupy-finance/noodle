package keeper

import (
	"sort"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	staking "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/soupy-finance/noodle/x/oracle/types"
)

type PriceInfo struct {
	val    sdk.Dec
	weight int64
}
type ValPriceList struct {
	prices    []PriceInfo
	netWeight int64
}

func (p ValPriceList) Len() int { return len(p.prices) }
func (p ValPriceList) Swap(i, j int) {
	temp := p.prices[j]
	p.prices[j] = p.prices[i]
	p.prices[i] = temp
}
func (p ValPriceList) Less(i, j int) bool { return p.prices[i].val.LT(p.prices[j].val) }

func (k Keeper) AggAssetPrice(ctx sdk.Context, asset string) string {
	var aggPrice sdk.Dec
	var priceList ValPriceList
	k.stakingKeeper.IterateBondedValidatorsByPower(ctx, k.GetValidatorPriceFn(ctx, asset, &priceList))

	// Process prices slice
	sort.Sort(priceList)
	var weightSeen int64 = 0

	for _, priceInfo := range priceList.prices {
		weightSeen += priceInfo.weight

		if weightSeen > priceList.netWeight {
			aggPrice = priceInfo.val
			break
		}
	}

	return aggPrice.String()
}

func (k Keeper) GetValidatorPriceFn(ctx sdk.Context, asset string, priceList *ValPriceList) (fn func(index int64, validator staking.ValidatorI) (stop bool)) {
	// Get store
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.PricesKey))

	return func(index int64, validator staking.ValidatorI) (stop bool) {
		// Get price from store and update prices slice
		valAssetKeyBytes := append(validator.GetOperator(), []byte(":"+asset)...)
		priceBytes := store.Get(valAssetKeyBytes)
		price, err := sdk.NewDecFromStr(string(priceBytes))

		if err != nil {
			panic(err)
		}

		weight := validator.GetConsensusPower(validator.GetBondedTokens())
		priceInfo := PriceInfo{price, weight}
		(*priceList).prices = append((*priceList).prices, priceInfo)
		(*priceList).netWeight += weight

		return false
	}
}
