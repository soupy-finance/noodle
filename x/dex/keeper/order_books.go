package keeper

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/soupy-finance/noodle/x/dex/types"
)

// Returns book without AMM orders
func (k Keeper) GetPureBook(ctx sdk.Context, market string, side byte) (*types.OrderBook, error) {
	// Validate arguments
	if side != 'b' && side != 's' {
		return nil, types.InvalidSide
	}

	markets := k.Markets(ctx)

	_, ok := markets[market]

	if !ok {
		return nil, types.InvalidMarket
	}

	book := types.OrderBook{
		Market: market,
		Side:   side,
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.BooksStoreKey))
	marketKeyBytes := []byte(market)
	storedBookBytes := store.Get(marketKeyBytes)
	var storedBook []types.StoredLevel

	if storedBookBytes != nil {
		err := json.Unmarshal(storedBookBytes, &storedBook)

		if err != nil {
			panic(err)
		}
	}

	// Allocate extra memory to make new order insertion more efficient
	book.Levels = make([]types.BookLevel, len(storedBook), len(storedBook)+1)

	for i, storedLevel := range storedBook {
		price, ok := sdk.NewIntFromString(storedLevel.Price)

		if !ok {
			panic("invalid stored price")
		}

		level := &book.Levels[i]
		*level = types.BookLevel{
			Market: market,
			Side:   side,
			Price:  price,
		}
		// Allocate extra memory to make new order insertion more efficient
		level.Orders = make([]types.Order, len(storedLevel.Orders), len(storedLevel.Orders)+1)

		for j, storedOrder := range storedLevel.Orders {
			order := &level.Orders[j]
			*order = types.Order{
				Account: storedOrder.Account,
				Market:  market,
				Side:    side,
				Price:   level.Price,
			}
			quantity, ok := sdk.NewIntFromString(storedOrder.Quantity)

			if !ok {
				panic("invalid stored quantity")
			}

			order.Quantity = quantity
		}
	}

	if len(book.Levels) > 0 {
		book.BestPrice = book.Levels[0].Price
	}

	return &book, nil
}

// Returns book with AMM orders
func (k Keeper) GetVirtualBook(ctx sdk.Context, market string, side byte) (*types.OrderBook, error) {
	book, err := k.GetPureBook(ctx, market, side)

	if err != nil {
		return nil, err
	}

	// Get AMM price and insert into book
	//

	return book, nil
}

func (k Keeper) SavePureBook(ctx sdk.Context, book *types.OrderBook) error {
	storedBook := make([]types.StoredLevel, len(book.Levels))

	for i, level := range book.Levels {
		storedLevel := &storedBook[i]
		*storedLevel = types.StoredLevel{
			Price:  level.Price.String(),
			Orders: make([]types.StoredOrder, len(level.Orders)),
		}

		for j, order := range level.Orders {
			storedOrder := &storedLevel.Orders[j]
			*storedOrder = types.StoredOrder{
				Account:  order.Account,
				Quantity: order.Quantity.String(),
			}
		}
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.BooksStoreKey))
	marketKeyBytes := []byte(book.Market)
	storedBookBytes, err := json.Marshal(storedBook)

	if err != nil {
		return err
	}

	store.Set(marketKeyBytes, storedBookBytes)
	return nil
}

func (k Keeper) InsertOrder(ctx sdk.Context, order *types.Order) error {
	book, err := k.GetPureBook(ctx, order.Market, order.Side)

	if err != nil {
		return err
	}

	inserted := false

	for i, level := range book.Levels {
		if level.Price.Equal(order.Price) {
			level.Orders = append(level.Orders, *order)
			inserted = true
			break
		} else if (order.Side == 'b' && level.Price.LT(order.Price)) ||
			(order.Side == 's' && level.Price.GT(order.Price)) {
			newLevel := types.BookLevel{
				Market: order.Market,
				Side:   order.Side,
				Price:  order.Price,
				Orders: []types.Order{*order},
			}

			book.Levels = append(book.Levels[:i+1], book.Levels[i:]...)
			book.Levels[i] = newLevel
			inserted = true
			break
		}
	}

	if !inserted {
		newLevel := types.BookLevel{
			Market: order.Market,
			Side:   order.Side,
			Price:  order.Price,
			Orders: []types.Order{*order},
		}

		book.Levels = append(book.Levels, newLevel)
	}

	err = k.SavePureBook(ctx, book)
	return err
}
