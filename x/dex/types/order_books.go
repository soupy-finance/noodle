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
	Market string
	Side   Side
	Levels []BookLevel
}

type BookLevel struct {
	Market string
	Side   Side
	Price  sdk.Dec
	Orders []Order
}

type Order struct {
	Account  sdk.AccAddress
	Market   string
	Side     Side
	Price    sdk.Dec
	Quantity sdk.Dec
}

type StoredLevel struct {
	Price  string
	Orders []StoredOrder
}

type StoredOrder struct {
	Account  string
	Quantity string
}

func NewSide(sideByte byte) (Side, bool) {
	side := Side(sideByte)
	return side, side.IsValid()
}

func (side *Side) IsValid() bool {
	switch *side {
	case Bid, Ask:
		return true
	default:
		return false
	}
}

func (side *Side) OppositeSide() Side {
	switch *side {
	case Bid:
		return 'a'
	case Ask:
		return 'b'
	default:
		return 'b'
	}
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

func (book *OrderBook) BestPrice() sdk.Dec {
	if len(book.Levels) > 0 {
		// Panics if level is empty
		return book.Levels[0].Price
	} else {
		return sdk.ZeroDec()
	}
}

func (book *OrderBook) RemoveTopLevel() {
	book.Levels = book.Levels[1:]
}
