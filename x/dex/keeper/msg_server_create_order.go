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

	orderType, ok := types.NewOrderType(msg.OrderType)

	if !ok {
		return nil, types.InvalidOrderType
	}

	markets := k.Markets(ctx)
	_, ok = markets[msg.Market]

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

	// Create order object and get order book
	order := types.Order{
		Account:  sdk.AccAddress(msg.Creator),
		Market:   msg.Market,
		Side:     side,
		Price:    price,
		Quantity: quantity,
	}
	book, err := k.Keeper.GetVirtualBook(ctx, msg.Market, side)

	if err != nil {
		return nil, err
	}

	// Process orders by type
	switch orderType {
	case types.MarketOrder:
		err = k.Keeper.ProcessMarketOrder(ctx, &order, book)
	case types.LimitOrder:
	case types.LimitFoKOrder:
	case types.LimitIoCOrder:
	case types.LimitPostOnlyOrder:
	}

	if err != nil {
		return nil, err
	}

	return &types.MsgCreateOrderResponse{}, nil
}

func (k Keeper) ProcessMarketOrder(ctx sdk.Context, order *types.Order, book *types.OrderBook) error {
	if len(book.Levels) == 0 {
		return types.NoOrdersInBook
	}

	return nil
}
