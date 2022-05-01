package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type OrderBook struct {
	Market    string
	Side      byte // b = buy, s = sell
	BestPrice sdk.Int
	Levels    []BookLevel
}

type BookLevel struct {
	Market string
	Side   byte
	Price  sdk.Int
	Orders []Order
}

type Order struct {
	Account  string
	Market   string
	Side     byte
	Price    sdk.Int
	Quantity sdk.Int
}

type StoredLevel struct {
	Price  string
	Orders []StoredOrder
}

type StoredOrder struct {
	Account  string
	Quantity string
}
