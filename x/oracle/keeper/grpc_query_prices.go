package keeper

import (
	"context"
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/soupy-finance/noodle/x/oracle/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Prices(goCtx context.Context, req *types.QueryPricesRequest) (*types.QueryPricesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	var assets []string
	err := json.Unmarshal([]byte(req.Assets), &assets)

	if err != nil {
		return nil, types.InvalidAssetList
	}

	prices := make([]uint64, len(assets))

	for i, asset := range assets {
		prices[i] = k.AggAssetPrice(ctx, asset)
	}

	pricesBytes, err := json.Marshal(prices)

	if err != nil {
		return nil, types.PriceAggError
	}

	return &types.QueryPricesResponse{Data: string(pricesBytes)}, nil
}
