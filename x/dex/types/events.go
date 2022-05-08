package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	EventTypeAddOffer      = "add_offer"
	EventTypeRemoveOffer   = "remove_offer"
	EventTypeTradeExec     = "trade_exec"
	EventTypeBalanceChange = "balance_change"

	AttributeKeyAccount  = "account"
	AttributeKeyMaker    = "maker"
	AttributeKeyTaker    = "taker"
	AttributeKeyMarket   = "market"
	AttributeKeyQuantity = "quantity"
	AttributeKeyPrice    = "price"
	AttributeKeySide     = "side"
	AttributeKeyAsset    = "asset"
)

func NewAddOfferEvent(account sdk.AccAddress, market string, quantity, price sdk.Dec, side Side) sdk.Event {
	return sdk.NewEvent(
		EventTypeAddOffer,
		sdk.NewAttribute(AttributeKeyAccount, account.String()),
		sdk.NewAttribute(AttributeKeyMarket, market),
		sdk.NewAttribute(AttributeKeyQuantity, quantity.String()),
		sdk.NewAttribute(AttributeKeyPrice, price.String()),
		sdk.NewAttribute(AttributeKeySide, string(side)),
	)
}

func NewRemoveOfferEvent(account sdk.AccAddress, market string, quantity, price sdk.Dec, side Side) sdk.Event {
	return sdk.NewEvent(
		EventTypeRemoveOffer,
		sdk.NewAttribute(AttributeKeyAccount, account.String()),
		sdk.NewAttribute(AttributeKeyMarket, market),
		sdk.NewAttribute(AttributeKeyQuantity, quantity.String()),
		sdk.NewAttribute(AttributeKeyPrice, price.String()),
		sdk.NewAttribute(AttributeKeySide, string(side)),
	)
}

func NewTradeExecEvent(maker, taker sdk.AccAddress, market string, quantity, price sdk.Dec) sdk.Event {
	return sdk.NewEvent(
		EventTypeRemoveOffer,
		sdk.NewAttribute(AttributeKeyMaker, maker.String()),
		sdk.NewAttribute(AttributeKeyTaker, taker.String()),
		sdk.NewAttribute(AttributeKeyMarket, market),
		sdk.NewAttribute(AttributeKeyQuantity, quantity.String()),
		sdk.NewAttribute(AttributeKeyPrice, price.String()),
	)
}

func NewBalanceChangeEvent(account sdk.AccAddress, asset string, quantity sdk.Dec) sdk.Event {
	return sdk.NewEvent(
		EventTypeRemoveOffer,
		sdk.NewAttribute(AttributeKeyAccount, account.String()),
		sdk.NewAttribute(AttributeKeyAsset, asset),
		sdk.NewAttribute(AttributeKeyQuantity, quantity.String()),
	)
}