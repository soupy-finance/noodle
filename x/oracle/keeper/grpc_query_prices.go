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
	prices := make([]string, len(req.Assets))
	prices[0] = k.AggAssetPrice(ctx, req.Assets[0])

	for i, asset := range req.Assets {
		prices[i] = k.AggAssetPrice(ctx, asset)
	}

	pricesBytes, err := json.Marshal(prices)

	if err != nil {
		panic(err)
	}

	return &types.QueryPricesResponse{Data: string(pricesBytes)}, nil
}
