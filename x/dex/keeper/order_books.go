package keeper

import (
	"encoding/binary"
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/soupy-finance/noodle/x/dex/types"
)

// Returns book without AMM orders
func (k Keeper) GetPureBook(ctx sdk.Context, market string, side types.Side) (*types.OrderBook, error) {
	book := types.OrderBook{
		Market: market,
		Side:   side,
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.BooksStoreKey))
	bookKeyBytes := []byte(market + ":" + string(side))
	storedBookBytes := store.Get(bookKeyBytes)
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
		price := sdk.MustNewDecFromStr(storedLevel.Price)

		level := &book.Levels[i]
		*level = types.BookLevel{
			Market: market,
			Side:   side,
			Price:  price,
		}
		// Allocate extra memory to make new order insertion more efficient
		level.Orders = make([]types.Order, len(storedLevel.Orders), len(storedLevel.Orders)+1)

		for j, storedOrder := range storedLevel.Orders {
			accAddress, err := sdk.AccAddressFromBech32(storedOrder.Account)

			if err != nil {
				panic(err)
			}

			order := &level.Orders[j]
			*order = types.Order{
				Id:      types.OrderId(storedOrder.Id),
				Account: accAddress,
				Market:  market,
				Side:    side,
				Price:   level.Price,
			}

			quantity := sdk.MustNewDecFromStr(storedOrder.Quantity)
			order.Quantity = quantity
		}
	}

	return &book, nil
}

// Returns book with AMM orders
func (k Keeper) GetVirtualBook(ctx sdk.Context, market string, side types.Side) (*types.OrderBook, error) {
	book, err := k.GetPureBook(ctx, market, side)

	if err != nil {
		return nil, err
	}

	// Get AMM price and insert into book
	//

	return book, nil
}

func (k Keeper) GetAccountOrders(ctx sdk.Context, account sdk.AccAddress) types.AccountOrders {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.AccountOrdersStoreKey))
	storedOrdersBytes := store.Get(account)
	storedOrders := map[types.OrderId]types.StoredAccountOrder{}

	if storedOrdersBytes != nil {
		err := json.Unmarshal(storedOrdersBytes, &storedOrders)

		if err != nil {
			panic(err)
		}
	}

	orders := types.AccountOrders{}

	for orderId, order := range storedOrders {
		side, ok := types.NewSide(order.Side[0])

		if !ok {
			panic("invalid stored account order")
		}

		orders[orderId] = types.Order{
			Id:       types.OrderId(orderId),
			Account:  account,
			Quantity: sdk.MustNewDecFromStr(order.Quantity),
			Price:    sdk.MustNewDecFromStr(order.Price),
			Side:     side,
		}
	}

	return orders
}

func (k Keeper) GetAccountOrdersCount(ctx sdk.Context, account sdk.AccAddress) (ordersCount uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.AccountOrdersCountKey))
	ordersCountBytes := store.Get(account)

	if ordersCountBytes == nil {
		ordersCount = 0
	} else {
		ordersCount = binary.BigEndian.Uint64(ordersCountBytes)
	}

	return
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
				Id:       string(order.Id),
				Account:  order.Account.String(),
				Quantity: order.Quantity.String(),
			}
		}
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.BooksStoreKey))
	bookKeyBytes := []byte(book.Market + ":" + string(book.Side))
	storedBookBytes, err := json.Marshal(storedBook)

	if err != nil {
		return err
	}

	store.Set(bookKeyBytes, storedBookBytes)
	return nil
}

func (k Keeper) SaveVirtualBook(ctx sdk.Context, book *types.OrderBook) error {
	// Remove AMM orders
	for i := range book.Levels {
		level := &book.Levels[i]

		for j, order := range level.Orders {
			if order.IsAmm {
				level.Orders = append(level.Orders[:j], level.Orders[j+1:]...)
				j--
			}
		}

		if len(level.Orders) == 0 {
			book.Levels = append(book.Levels[:i], book.Levels[i+1:]...)
			i--
		}
	}

	if len(book.Levels) == 0 {
		store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.BooksStoreKey))
		bookKeyBytes := []byte(book.Market + ":" + string(book.Side))
		store.Delete(bookKeyBytes)
		return nil
	}

	err := k.SavePureBook(ctx, book)
	return err
}

func (k Keeper) InsertAccountOrder(ctx sdk.Context, order *types.Order) error {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.AccountOrdersStoreKey))
	accountOrdersBytes := store.Get(order.Account)
	accountOrders := map[types.OrderId]types.StoredAccountOrder{}

	if accountOrdersBytes != nil {
		err := json.Unmarshal(accountOrdersBytes, &accountOrders)

		if err != nil {
			panic(err)
		}
	}

	accountOrders[order.Id] = types.StoredAccountOrder{
		Market:   order.Market,
		Price:    order.Price.String(),
		Quantity: order.Quantity.String(),
		Side:     string(order.Side),
		Filled:   "0",
		Date:     ctx.BlockTime().Unix(),
	}

	accountOrdersBytes, err := json.Marshal(accountOrders)

	if err != nil {
		return err
	}

	store.Set(order.Account, accountOrdersBytes)
	k.IncrementAccountOrdersCount(ctx, order.Account)
	return nil
}

func (k Keeper) UpdateAccountOrder(ctx sdk.Context, order *types.Order, quantityChange sdk.Dec) error {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.AccountOrdersStoreKey))
	accountOrdersBytes := store.Get(order.Account)
	accountOrders := map[types.OrderId]types.StoredAccountOrder{}

	if accountOrdersBytes == nil {
		panic("updating nonexistent order")
	}

	err := json.Unmarshal(accountOrdersBytes, &accountOrders)

	if err != nil {
		panic(err)
	}

	accountOrder, exists := accountOrders[order.Id]

	if !exists {
		panic("updating nonexistent order")
	}

	accountOrder.Filled = sdk.MustNewDecFromStr(accountOrder.Filled).Add(quantityChange).String()
	accountOrders[order.Id] = accountOrder
	accountOrdersBytes, err = json.Marshal(accountOrders)

	if err != nil {
		return err
	}

	store.Set(order.Account, accountOrdersBytes)
	k.IncrementAccountOrdersCount(ctx, order.Account)
	return nil
}

func (k Keeper) RemoveAccountOrder(ctx sdk.Context, order *types.Order) error {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.AccountOrdersStoreKey))
	accountOrdersBytes := store.Get(order.Account)
	accountOrders := map[types.OrderId]types.StoredAccountOrder{}

	if accountOrdersBytes != nil {
		err := json.Unmarshal(accountOrdersBytes, &accountOrders)

		if err != nil {
			panic(err)
		}

		delete(accountOrders, order.Id)
		accountOrdersBytes, err := json.Marshal(accountOrders)

		if err != nil {
			return err
		}

		store.Set(order.Account, accountOrdersBytes)
	}

	return nil
}

func (k Keeper) InsertOrder(ctx sdk.Context, order *types.Order, book *types.OrderBook, save bool) error {
	var err error

	if book == nil {
		book, err = k.GetPureBook(ctx, order.Market, order.Side)
	}

	if err != nil {
		return err
	}

	inserted := false

	for i := range book.Levels {
		level := &book.Levels[i]

		if level.Price.Equal(order.Price) {
			level.Orders = append(level.Orders, *order)
			inserted = true
			break
		} else if (order.Side == 'b' && level.Price.LT(order.Price)) ||
			(order.Side == 'a' && level.Price.GT(order.Price)) {
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

	if save {
		err = k.SaveVirtualBook(ctx, book)

		if err != nil {
			return err
		}
	}

	k.InsertAccountOrder(ctx, order)
	return nil
}

func (k Keeper) RemoveOrder(ctx sdk.Context, order *types.Order, book *types.OrderBook, save bool) error {
	var err error

	if book == nil {
		book, err = k.GetPureBook(ctx, order.Market, order.Side)
	}

	if err != nil {
		return err
	}

	for i := range book.Levels {
		level := &book.Levels[i]

		if level.Price.Equal(order.Price) {
			for j := range level.Orders {
				if level.Orders[j].Id == order.Id {
					level.Orders = append(level.Orders[:j], level.Orders[j+1:]...)
					break
				}
			}
		}
	}

	if save {
		err = k.SaveVirtualBook(ctx, book)

		if err != nil {
			return err
		}
	}

	k.RemoveAccountOrder(ctx, order)
	return nil
}

func (k Keeper) IncrementAccountOrdersCount(ctx sdk.Context, account sdk.AccAddress) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.AccountOrdersCountKey))
	ordersCountBytes := store.Get(account)
	var ordersCount uint64

	if ordersCountBytes == nil {
		ordersCount = 0
		ordersCountBytes = make([]byte, 8)
	} else {
		ordersCount = binary.BigEndian.Uint64(ordersCountBytes)
	}

	binary.BigEndian.PutUint64(ordersCountBytes, ordersCount+1)
	store.Set(account, ordersCountBytes)
}

func (k Keeper) SaveBooks(ctx sdk.Context) {
	// store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.BooksStoreKey))
	// bookKeyBytes := []byte(market + ":" + string(side))
	// storedBookBytes := store.Get(bookKeyBytes)
}
