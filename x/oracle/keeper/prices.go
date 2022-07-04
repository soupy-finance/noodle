package keeper

import (
	"encoding/json"
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

func (k Keeper) GetAssetPrices(ctx sdk.Context, assets []string) []string {
	prices := make([]string, len(assets))

	for i, asset := range assets {
		price, exists := k.prices[asset]

		if !exists {
			price = sdk.ZeroDec()
		}

		prices[i] = price.String()
	}

	return prices
}

func (k Keeper) AggAssetPrice(ctx sdk.Context, asset string) sdk.Dec {
	var priceList ValPriceList
	aggPrice := sdk.ZeroDec()

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

	return aggPrice
}

func (k Keeper) GetValidatorPriceFn(ctx sdk.Context, asset string, priceList *ValPriceList) (fn func(index int64, validator staking.ValidatorI) (stop bool)) {
	return func(index int64, validator staking.ValidatorI) (stop bool) {
		price, exists := k.valPrices[asset]

		if !exists {
			return false
		}

		weight := validator.GetConsensusPower(validator.GetBondedTokens())
		priceInfo := PriceInfo{price, weight}
		(*priceList).prices = append((*priceList).prices, priceInfo)
		(*priceList).netWeight += weight

		return false
	}
}

func (k Keeper) CachePrices(ctx sdk.Context) {
	if k.pricesCached {
		return
	}

	assets := k.AssetsParsed(ctx)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.PricesKey))

	for _, asset := range assets {
		priceStr := store.Get([]byte(asset))

		if priceStr != nil {
			k.prices[asset] = sdk.MustNewDecFromStr(string(priceStr))
		}
	}

	k.pricesCached = true
}

func (k Keeper) UpdateAndSavePrices(ctx sdk.Context) {
	assets := k.AssetsParsed(ctx)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.PricesKey))

	for _, asset := range assets {
		price := k.AggAssetPrice(ctx, asset)
		k.prices[asset] = price
		store.Set([]byte(asset), []byte(price.String()))
	}

	pricesStr, err := json.Marshal(k.prices)

	if err != nil {
		panic(err)
	}

	EmitPricesEvents(ctx, string(pricesStr))
}
