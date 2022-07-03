package types

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type OrderId string
type Side byte
type OrderType string
type OrderFlag string
type OrderFlags map[OrderFlag]bool

const (
	Bid Side = 'b'
	Ask Side = 'a'
)

const (
	MarketOrder OrderType = "market"
	LimitOrder  OrderType = "limit"
)

const (
	NoFlag       OrderFlag = "none"
	FoKFlag      OrderFlag = "fill_or_kill"
	IoCFlag      OrderFlag = "immediate_or_cancel"
	PostOnlyFlag OrderFlag = "post_only"
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
	Id       OrderId
	Account  sdk.AccAddress
	Market   string
	Side     Side
	Price    sdk.Dec
	Quantity sdk.Dec
	Type     OrderType
	Flags    OrderFlags
	IsAmm    bool
}

type AccountOrders map[OrderId]Order

type StoredLevel struct {
	Price  string
	Orders []StoredOrder
}

type StoredOrder struct {
	Id       string
	Account  string
	Quantity string
}

type StoredAccountOrder struct {
	Market   string
	Price    string
	Quantity string
	Filled   string
	Side     string
	Date     int64
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
	case MarketOrder, LimitOrder:
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

func NewOrderFlag(flagStr string) (OrderFlag, bool) {
	flag := OrderFlag(flagStr)
	return flag, flag.IsValid()
}

func (flag *OrderFlag) IsValid() bool {
	switch *flag {
	case NoFlag, FoKFlag, IoCFlag, PostOnlyFlag:
		return true
	default:
		return false
	}
}

func NewOrderFlags(flags []string) (OrderFlags, bool) {
	orderFlags := OrderFlags{}
	ok := true

	for _, flagStr := range flags {
		flag, _ok := NewOrderFlag(flagStr)
		ok = ok && _ok
		orderFlags[flag] = true
	}

	return orderFlags, ok
}

func NewOrderId(account sdk.AccAddress, ordersCount uint64) OrderId {
	idBytes := make([]byte, 16) // 8 bytes for address, 8 for count
	copy(idBytes, account)
	binary.BigEndian.PutUint64(idBytes[len(idBytes)-8:], ordersCount)
	hash := sha256.Sum256(idBytes)
	return OrderId(hex.EncodeToString(hash[:]))
}
