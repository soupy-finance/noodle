package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	EventTypePrices = "prices"

	AttributeKeyData = "data"
)

func NewPricesEvent(prices string) sdk.Event {
	return sdk.NewEvent(
		EventTypePrices,
		sdk.NewAttribute(AttributeKeyData, prices),
	)
}
