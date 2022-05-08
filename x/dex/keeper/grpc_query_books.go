package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/soupy-finance/noodle/x/dex/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Books(goCtx context.Context, req *types.QueryBooksRequest) (*types.QueryBooksResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate market
	markets := k.Markets(ctx)
	_, ok := markets[req.Market]

	if !ok {
		return nil, types.InvalidMarket
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.BooksStoreKey))
	bidsKeyBytes := []byte(req.Market + ":b")
	asksKeyBytes := []byte(req.Market + ":a")
	bidsBytes := store.Get(bidsKeyBytes)
	asksBytes := store.Get(asksKeyBytes)

	// Get AMM data
	//
	ammStr := ""

	return &types.QueryBooksResponse{
		Bids: string(bidsBytes),
		Asks: string(asksBytes),
		Amm:  ammStr,
	}, nil
}
