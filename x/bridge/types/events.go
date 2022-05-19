package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	EventTypeWithdraw = "withdraw"

	AttributeKeyWId      = "withdrawal_id"
	AttributeKeyAsset    = "asset"
	AttributeKeyQuantity = "quantity"
	AttributeKeyAddress  = "address"
	AttributeKeyChainId  = "chain_id"
)

func NewWithdrawEvent(wId string, asset string, quantity string, address string, chainId string) sdk.Event {
	return sdk.NewEvent(
		EventTypeWithdraw,
		sdk.NewAttribute(AttributeKeyWId, wId),
		sdk.NewAttribute(AttributeKeyAsset, asset),
		sdk.NewAttribute(AttributeKeyQuantity, quantity),
		sdk.NewAttribute(AttributeKeyAddress, address),
		sdk.NewAttribute(AttributeKeyChainId, chainId),
	)
}
