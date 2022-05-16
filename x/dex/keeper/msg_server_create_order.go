package keeper

import (
	"context"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/soupy-finance/noodle/x/dex/types"
)

var zeroDec sdk.Dec = sdk.ZeroDec()

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

	markets := k.MarketsParsed(ctx)
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

	flags, ok := types.NewOrderFlags(msg.Flags)

	if !ok {
		return nil, types.InvalidOrderFlags
	}

	// Create order id from account address and historical order count
	account := sdk.AccAddress(msg.Creator)
	accountOrdersCount := k.GetAccountOrdersCount(ctx, account)
	orderId := types.NewOrderId(account, accountOrdersCount)

	// Create order object and get order book
	order := types.Order{
		Id:       orderId,
		Account:  account,
		Market:   msg.Market,
		Side:     side,
		Price:    price,
		Quantity: quantity,
		Type:     orderType,
		Flags:    flags,
	}
	bids, err := k.Keeper.GetVirtualBook(ctx, msg.Market, types.Bid)

	if err != nil {
		return nil, err
	}

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
	}

	if err != nil {
		return nil, err
	}

	err = k.SaveVirtualBook(ctx, bids)

	if err != nil {
		return nil, err
	}

	err = k.SaveVirtualBook(ctx, asks)

	if err != nil {
		return nil, err
	}

	return &types.MsgCreateOrderResponse{}, nil
}

func (k Keeper) ProcessMarketOrder(
	ctx sdk.Context,
	order *types.Order,
	bids *types.OrderBook,
	asks *types.OrderBook,
) error {
	var book *types.OrderBook

	switch order.Side {
	case types.Bid:
		book = asks
	case types.Ask:
		book = bids
	}

	// Match orders
	execSent, execRecv, sendAsset, recvAsset := k.MatchUntilEmpty(ctx, order, book)

	if order.Quantity.GT(zeroDec) {
		return types.InsufficientLiquidity
	}

	EmitBalanceChangeEvent(ctx, order.Account, sendAsset, execSent.Neg())
	EmitBalanceChangeEvent(ctx, order.Account, recvAsset, execRecv)
	return nil
}

func (k Keeper) ProcessLimitOrder(
	ctx sdk.Context,
	order *types.Order,
	bids *types.OrderBook,
	asks *types.OrderBook,
) error {
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

	// Process post-only flag
	//

	execSent, execRecv, sendAsset, recvAsset := k.MatchUntilEmpty(ctx, order, fillBook)

	// Process other flags
	//

	// Insert remaining part of order into book
	if order.Quantity.GT(zeroDec) {
		err := k.InsertOrder(ctx, order, insertBook, false)

		if err != nil {
			return err
		}

		// Store tokens in escrow
		makerSendCoins := sdk.NewCoins(sdk.NewCoin(sendAsset, sdk.NewIntFromBigInt(order.Quantity.BigInt())))
		k.bankKeeper.SendCoinsFromAccountToModule(ctx, order.Account, types.ModuleName, makerSendCoins)

		if err != nil {
			return err
		}

		EmitAddOfferEvent(ctx, order.Account, order.Id, order.Market, order.Quantity, order.Price, order.Side)
	}

	EmitBalanceChangeEvent(ctx, order.Account, sendAsset, execSent.Neg())
	EmitBalanceChangeEvent(ctx, order.Account, recvAsset, execRecv)
	return nil
}

func (k Keeper) SufficientBalanceForLimitOrder(ctx sdk.Context, order *types.Order) bool {
	var asset string
	var quantity sdk.Dec
	assets := strings.Split(order.Market, "-")

	switch order.Side {
	case types.Bid:
		asset = assets[1]
		quantity = order.Quantity.Mul(order.Price)
	case types.Ask:
		asset = assets[0]
		quantity = order.Quantity
	}

	return k.SufficientBalanceForQuantity(ctx, order.Account, quantity, asset)
}

func (k Keeper) SufficientBalanceForQuantity(
	ctx sdk.Context,
	account sdk.AccAddress,
	quantity sdk.Dec,
	asset string,
) bool {
	balance := k.bankKeeper.GetBalance(ctx, account, asset)
	quantityInt := sdk.NewIntFromBigInt(quantity.BigInt())

	return balance.Amount.GTE(quantityInt)
}

func (k Keeper) MatchUntilEmpty(
	ctx sdk.Context,
	order *types.Order,
	book *types.OrderBook,
) (execSent, execRecv sdk.Dec, sendAsset, recvAsset string) {
	assets := strings.Split(order.Market, "-")
	execRecv = zeroDec
	execSent = zeroDec

	if order.Side == types.Ask {
		sendAsset = assets[0]
		recvAsset = assets[1]

		for len(book.Levels) > 0 &&
			order.Quantity.GT(zeroDec) &&
			(order.Type == types.MarketOrder || order.Price.LTE(book.BestPrice())) {

			execSent, execRecv = k.MatchNextBestBid(ctx, order, book, sendAsset, recvAsset, execSent, execRecv)
		}
	} else {
		sendAsset = assets[1]
		recvAsset = assets[0]

		for len(book.Levels) > 0 &&
			order.Quantity.GT(zeroDec) &&
			(order.Type == types.MarketOrder || order.Price.GTE(book.BestPrice())) {

			execSent, execRecv = k.MatchNextBestAsk(ctx, order, book, sendAsset, recvAsset, execSent, execRecv)
		}
	}

	return
}

func (k Keeper) MatchNextBestAsk(
	ctx sdk.Context,
	order *types.Order,
	book *types.OrderBook,
	sendAsset string,
	recvAsset string,
	execSent sdk.Dec,
	execRecv sdk.Dec,
) (sdk.Dec, sdk.Dec) {
	level := book.Levels[0]
	orders := level.Orders
	// Panics if the level is empty
	bestOffer := level.Orders[0]

	localExecRecv := sdk.MinDec(order.Quantity, bestOffer.Quantity)
	localExecSent := localExecRecv.Mul(bestOffer.Price)
	execSent = execSent.Add(localExecSent)
	execRecv = execRecv.Add(localExecRecv)
	order.Quantity = order.Quantity.Sub(localExecRecv)
	bestOffer.Quantity = bestOffer.Quantity.Sub(localExecRecv)

	// Transfer asset directly from taker to maker
	takerSendCoins := sdk.NewCoins(sdk.NewCoin(sendAsset, sdk.NewIntFromBigInt(localExecSent.BigInt())))
	k.bankKeeper.SendCoins(ctx, order.Account, bestOffer.Account, takerSendCoins)

	// Transfer asset from escrow to taker
	takerRecvCoins := sdk.NewCoins(sdk.NewCoin(recvAsset, sdk.NewIntFromBigInt(localExecRecv.BigInt())))
	k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, order.Account, takerRecvCoins)

	// Remove price level if empty
	if bestOffer.Quantity.LTE(zeroDec) {
		level.Orders = orders[1:]

		if len(level.Orders) == 0 {
			book.RemoveTopLevel()
		}

		EmitRemoveOfferEvent(ctx, bestOffer.Account, bestOffer.Id)
	} else {
		EmitUpdateOfferEvent(ctx, bestOffer.Account, bestOffer.Id, bestOffer.Quantity)
	}

	EmitTradeExecEvent(ctx, bestOffer.Account, order.Account, order.Market, localExecRecv, bestOffer.Price)
	return execSent, execRecv
}

func (k Keeper) MatchNextBestBid(
	ctx sdk.Context,
	order *types.Order,
	book *types.OrderBook,
	sendAsset string,
	recvAsset string,
	execSent sdk.Dec,
	execRecv sdk.Dec,
) (sdk.Dec, sdk.Dec) {
	level := book.Levels[0]
	orders := level.Orders
	// Panics if the level is empty
	bestOffer := level.Orders[0]

	localExecSent := sdk.MinDec(order.Quantity, bestOffer.Quantity)
	localExecRecv := localExecSent.Mul(bestOffer.Price)
	execSent = execSent.Add(localExecSent)
	execRecv = execRecv.Add(localExecRecv)
	order.Quantity = order.Quantity.Sub(localExecSent)
	bestOffer.Quantity = bestOffer.Quantity.Sub(localExecSent)

	// Transfer asset directly from taker to maker
	takerSendCoins := sdk.NewCoins(sdk.NewCoin(sendAsset, sdk.NewIntFromBigInt(localExecSent.BigInt())))
	k.bankKeeper.SendCoins(ctx, order.Account, bestOffer.Account, takerSendCoins)

	// Transfer asset from escrow to taker
	takerRecvCoins := sdk.NewCoins(sdk.NewCoin(recvAsset, sdk.NewIntFromBigInt(localExecRecv.BigInt())))
	k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, order.Account, takerRecvCoins)

	// Remove price level if empty
	if bestOffer.Quantity.LTE(zeroDec) {
		level.Orders = orders[1:]

		if len(level.Orders) == 0 {
			book.RemoveTopLevel()
		}
	}

	EmitTradeExecEvent(ctx, bestOffer.Account, order.Account, order.Market, localExecSent, bestOffer.Price)
	return execSent, execRecv
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

func GetSendAndRecvAssets(order *types.Order) (sendAsset, recvAsset string) {
	assets := strings.Split(order.Market, "-")

	if order.Side == types.Ask {
		sendAsset = assets[0]
		recvAsset = assets[1]
	} else {
		sendAsset = assets[1]
		recvAsset = assets[0]
	}

	return
}
