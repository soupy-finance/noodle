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

	price, err := sdk.NewDecFromStr(msg.Price)

	if err != nil {
		return nil, err
	}

	quantity, err := sdk.NewDecFromStr(msg.Quantity)

	if err != nil {
		return nil, err
	}

	// Create order object and get order book
	order := types.Order{
		Account:  sdk.AccAddress(msg.Creator),
		Market:   msg.Market,
		Side:     side,
		Price:    price,
		Quantity: quantity,
	}
	bids, err := k.Keeper.GetVirtualBook(ctx, msg.Market, types.Bid)
	asks, err := k.Keeper.GetVirtualBook(ctx, msg.Market, types.Ask)

	if err != nil {
		return nil, err
	}

	// Process orders by type
	switch orderType {
	case types.MarketOrder:
		err = k.Keeper.ProcessMarketOrder(ctx, &order, bids, asks)
	case types.LimitOrder:
		err = k.Keeper.ProcessLimitOrder(ctx, &order, bids, asks)
	case types.LimitFoKOrder:
	case types.LimitIoCOrder:
	case types.LimitPostOnlyOrder:
	}

	if err != nil {
		return nil, err
	}

	return &types.MsgCreateOrderResponse{}, nil
}

func (k Keeper) ProcessMarketOrder(ctx sdk.Context, order *types.Order, bids *types.OrderBook, asks *types.OrderBook) error {
	var book *types.OrderBook

	switch order.Side {
	case types.Bid:
		book = asks
	case types.Ask:
		book = bids
	}

	if len(book.Levels) == 0 {
		return types.NoOrdersInBook
	}

	// Match orders
	//

	return nil
}

func (k Keeper) ProcessLimitOrder(ctx sdk.Context, order *types.Order, bids *types.OrderBook, asks *types.OrderBook) error {
	if !SufficientBalanceForLimitOrder(order) {
		return types.InsufficientBalance
	}

	var fillBook *types.OrderBook
	var insertBook *types.OrderBook

	switch order.Side {
	case types.Bid:
		fillBook = asks
		insertBook = bids
	case types.Ask:
		fillBook = bids
		insertBook = asks
	}

	zero := sdk.ZeroDec()
	executedQuantity := zero
	executedCost := zero

	// Match existing orders
	for len(fillBook.Levels) > 0 && OrderIsLimitMatched(order, fillBook) && order.Quantity.GT(zero) {
		level := fillBook.Levels[0]
		orders := level.Orders
		// Panics if the level is empty
		bestOffer := level.Orders[0]

		localExecQuantity := sdk.MinDec(order.Quantity, bestOffer.Quantity)
		localExecCost := localExecQuantity.Mul(bestOffer.Price)
		executedQuantity = executedQuantity.Add(localExecQuantity)
		executedCost = executedCost.Add(localExecCost)
		order.Quantity = order.Quantity.Sub(localExecQuantity)
		bestOffer.Quantity = bestOffer.Quantity.Sub(localExecQuantity)

		if bestOffer.Quantity.LTE(zero) {
			level.Orders = orders[1:]

			if len(level.Orders) == 0 {
				fillBook.RemoveTopLevel()
			}
		}
	}

	// Insert remaining part of order into book
	if order.Quantity.GT(zero) {

	}

	return nil
}

func OrderIsLimitMatched(order *types.Order, book *types.OrderBook) bool {
	switch order.Side {
	case types.Bid:
		return order.Price.GTE(book.BestPrice())
	case types.Ask:
		return order.Price.LTE(book.BestPrice())
	default:
		return false
	}
}

func SufficientBalanceForLimitOrder(order *types.Order) bool {
	// Implement
	//

	return true
}
