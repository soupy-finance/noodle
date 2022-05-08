package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/soupy-finance/noodle/x/dex/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Book(goCtx context.Context, req *types.QueryBookRequest) (*types.QueryBookResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate arguments
	var side string

	if !req.Side {
		side = "b"
	} else {
		side = "a"
	}

	markets := k.Markets(ctx)
	_, ok := markets[req.Market]

	if !ok {
		return nil, types.InvalidMarket
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.BooksStoreKey))
	bookKeyBytes := []byte(req.Market + ":" + side)
	bookBytes := store.Get(bookKeyBytes)

	// Get AMM data
	//
	ammStr := ""

	return &types.QueryBookResponse{
		Pure: string(bookBytes),
		Amm:  ammStr,
	}, nil
}
