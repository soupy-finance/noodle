package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	EventTypeAddOffer      = "add_offer"
	EventTypeRemoveOffer   = "remove_offer"
	EventTypeUpdateOffer   = "update_offer"
	EventTypeTradeExec     = "trade_exec"
	EventTypeBalanceChange = "balance_change"

	AttributeKeyAccount  = "account"
	AttributeKeyMaker    = "maker"
	AttributeKeyTaker    = "taker"
	AttributeKeyMarket   = "market"
	AttributeKeyQuantity = "quantity"
	AttributeKeyPrice    = "price"
	AttributeKeySide     = "side"
	AttributeKeyOrderId  = "order_id"
	AttributeKeyAsset    = "asset"
)

func NewAddOfferEvent(id OrderId, account sdk.AccAddress, market string, quantity, price sdk.Dec, side Side) sdk.Event {
	return sdk.NewEvent(
		EventTypeAddOffer,
		sdk.NewAttribute(AttributeKeyOrderId, string(id)),
		sdk.NewAttribute(AttributeKeyAccount, toEventAccAddress(account)),
		sdk.NewAttribute(AttributeKeyMarket, market),
		sdk.NewAttribute(AttributeKeyQuantity, quantity.String()),
		sdk.NewAttribute(AttributeKeyPrice, price.String()),
		sdk.NewAttribute(AttributeKeySide, string(side)),
	)
}

func NewRemoveOfferEvent(id OrderId, account sdk.AccAddress, market string) sdk.Event {
	return sdk.NewEvent(
		EventTypeRemoveOffer,
		sdk.NewAttribute(AttributeKeyOrderId, string(id)),
		sdk.NewAttribute(AttributeKeyAccount, toEventAccAddress(account)),
		sdk.NewAttribute(AttributeKeyMarket, market),
	)
}

func NewUpdateOfferEvent(id OrderId, account sdk.AccAddress, market string, quantity sdk.Dec) sdk.Event {
	return sdk.NewEvent(
		EventTypeUpdateOffer,
		sdk.NewAttribute(AttributeKeyOrderId, string(id)),
		sdk.NewAttribute(AttributeKeyAccount, toEventAccAddress(account)),
		sdk.NewAttribute(AttributeKeyMarket, market),
		sdk.NewAttribute(AttributeKeyQuantity, quantity.String()),
	)
}

func NewTradeExecEvent(maker, taker sdk.AccAddress, market string, quantity, price sdk.Dec) sdk.Event {
	return sdk.NewEvent(
		EventTypeTradeExec,
		sdk.NewAttribute(AttributeKeyMaker, maker.String()),
		sdk.NewAttribute(AttributeKeyTaker, taker.String()),
		sdk.NewAttribute(AttributeKeyMarket, market),
		sdk.NewAttribute(AttributeKeyQuantity, quantity.String()),
		sdk.NewAttribute(AttributeKeyPrice, price.String()),
	)
}

func NewBalanceChangeEvent(account sdk.AccAddress, asset string, quantity sdk.Dec) sdk.Event {
	return sdk.NewEvent(
		EventTypeBalanceChange,
		sdk.NewAttribute(AttributeKeyAccount, toEventAccAddress(account)),
		sdk.NewAttribute(AttributeKeyAsset, asset),
		sdk.NewAttribute(AttributeKeyQuantity, quantity.String()),
	)
}

func toEventAccAddress(address sdk.AccAddress) string {
	prefix := sdk.GetConfig().GetBech32AccountAddrPrefix()
	return sdk.MustBech32ifyAddressBytes(prefix, address)
}
