package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/soupy-finance/noodle/x/dex/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) OpenOrders(goCtx context.Context, req *types.QueryOpenOrdersRequest) (*types.QueryOpenOrdersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	account, err := sdk.AccAddressFromBech32(req.Account)

	if err != nil {
		return nil, types.InvalidQueryAddress
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.AccountOrdersStoreKey))
	ordersBytes := store.Get(account)
	orders := string(ordersBytes)

	return &types.QueryOpenOrdersResponse{Orders: orders}, nil
}
