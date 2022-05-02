package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Side byte
type OrderType string

const (
	Bid Side = 'b'
	Ask Side = 'a'
)

const (
	MarketOrder        OrderType = "market"
	LimitOrder         OrderType = "limit"
	LimitFoKOrder      OrderType = "limit_fok"
	LimitIoCOrder      OrderType = "limit_ioc"
	LimitPostOnlyOrder OrderType = "limit_post_only"
)

type OrderBook struct {
	Market    string
	Side      Side
	BestPrice sdk.Int
	Levels    []BookLevel
}

type BookLevel struct {
	Market string
	Side   Side
	Price  sdk.Int
	Orders []Order
}

type Order struct {
	Account  sdk.AccAddress
	Market   string
	Side     Side
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

func NewOrderType(orderTypeStr string) (OrderType, bool) {
	orderType := OrderType(orderTypeStr)
	return orderType, orderType.IsValid()
}

func (t *OrderType) IsValid() bool {
	switch *t {
	case MarketOrder, LimitOrder, LimitFoKOrder, LimitIoCOrder, LimitPostOnlyOrder:
		return true
	default:
		return false
	}
}
