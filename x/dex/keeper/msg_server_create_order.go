package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/soupy-finance/noodle/x/dex/types"
)

func (k msgServer) CreateOrder(goCtx context.Context, msg *types.MsgCreateOrder) (*types.MsgCreateOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate arguments
	var side types.Side

	if !msg.Side {
		side = 'b'
	} else {
		side = 'a'
	}

	orderType := types.OrderType(msg.OrderType)

	if !orderType.IsValid() {
		return nil, types.InvalidOrderType
	}

	markets := k.Markets(ctx)
	_, ok := markets[msg.Market]

	if !ok {
		return nil, types.InvalidMarket
	}

	price, ok := sdk.NewIntFromString(msg.Price)

	if !ok {
		return nil, types.InvalidPrice
	}

	quantity, ok := sdk.NewIntFromString(msg.Quantity)

	if !ok {
		return nil, types.InvalidQuantity
	}

	order := types.Order{
		Account:  sdk.AccAddress(msg.Creator),
		Market:   msg.Market,
		Side:     side,
		Price:    price,
		Quantity: quantity,
	}

	switch orderType {
	case types.MarketOrder:

	}

	return &types.MsgCreateOrderResponse{}, nil
}
