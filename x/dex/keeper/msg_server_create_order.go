package keeper

import (
	"context"
	"strings"

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
	if !k.SufficientBalanceForLimitOrder(ctx, order) {
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
	assets := strings.Split(order.Market, "/")

	var sendAsset string
	var recvAsset string

	if order.Side == types.Ask {
		sendAsset = assets[0]
		recvAsset = assets[1]
	} else {
		sendAsset = assets[1]
		recvAsset = assets[0]
	}

	// Match existing orders
	for len(fillBook.Levels) > 0 && LimitOrderIsMatched(order, fillBook) && order.Quantity.GT(zero) {
		level := fillBook.Levels[0]
		orders := level.Orders
		// Panics if the level is empty
		bestOffer := level.Orders[0]

		localExecQuantity := sdk.MinDec(order.Quantity, bestOffer.Quantity)
		localExecQuote := localExecQuantity.Mul(bestOffer.Price)
		executedQuantity = executedQuantity.Add(localExecQuantity)
		executedCost = executedCost.Add(localExecQuote)
		order.Quantity = order.Quantity.Sub(localExecQuantity)
		bestOffer.Quantity = bestOffer.Quantity.Sub(localExecQuantity)

		// Transfer asset directly from taker to maker
		takerSendCoins := sdk.NewCoins(sdk.NewCoin(sendAsset, sdk.NewIntFromBigInt(localExecQuote.BigInt())))
		k.bankKeeper.SendCoins(ctx, order.Account, bestOffer.Account, takerSendCoins)

		// Transfer main asset from esrow to taker
		takerRecvCoins := sdk.NewCoins(sdk.NewCoin(recvAsset, sdk.NewIntFromBigInt(localExecQuantity.BigInt())))
		k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, order.Account, takerRecvCoins)

		// Remove price level if empty
		if bestOffer.Quantity.LTE(zero) {
			level.Orders = orders[1:]

			if len(level.Orders) == 0 {
				fillBook.RemoveTopLevel()
			}
		}
	}

	// Insert remaining part of order into book
	if order.Quantity.GT(zero) {
		inserted := false

		for i := 0; i < len(insertBook.Levels); i++ {
			level := insertBook.Levels[i]
			at, past := AtOrPastInsertionPoint(order, &level)

			if at {
				level.Orders = append(level.Orders, *order)
				break
			} else if past {
				// Insert new level at index
				// Include order in level's orders
				break
			}
		}

		if !inserted {
			// Apppend new level at end
			// Include order in level's orders
		}

		// Store tokens in escrow
		makerSendCoins := sdk.NewCoins(sdk.NewCoin(sendAsset, sdk.NewIntFromBigInt(order.Quantity.BigInt())))
		k.bankKeeper.SendCoinsFromAccountToModule(ctx, order.Account, types.ModuleName, makerSendCoins)
	}

	return nil
}

func (k Keeper) SufficientBalanceForLimitOrder(ctx sdk.Context, order *types.Order) bool {
	// Panics if market is not in a valid format
	quoteAsset := strings.Split(order.Market, "/")[1]
	balance := k.bankKeeper.GetBalance(ctx, order.Account, quoteAsset)
	quantityInt := sdk.NewIntFromBigInt(order.Quantity.BigInt())

	return balance.Amount.GTE(quantityInt)
}

func LimitOrderIsMatched(order *types.Order, book *types.OrderBook) bool {
	switch order.Side {
	case types.Bid:
		return order.Price.GTE(book.BestPrice())
	case types.Ask:
		return order.Price.LTE(book.BestPrice())
	default:
		return false
	}
}

func AtOrPastInsertionPoint(order *types.Order, level *types.BookLevel) (at, past bool) {
	switch order.Side {
	case types.Bid:
		return order.Price.GT(level.Price), order.Price.Equal(level.Price)
	case types.Ask:
		return order.Price.LT(level.Price), order.Price.Equal(level.Price)
	default:
		return false, false
	}
}
