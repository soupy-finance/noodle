package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	EventTypeWithdraw = "withdraw"

	AttributeKeyAccount  = "account"
	AttributeKeyWId      = "withdrawal_id"
	AttributeKeyAsset    = "asset"
	AttributeKeyQuantity = "quantity"
	AttributeKeyChainId  = "chain_id"
)

func NewWithdrawEvent(account string, wId string, asset string, quantity string, chainId string) sdk.Event {
	return sdk.NewEvent(
		EventTypeWithdraw,
		sdk.NewAttribute(AttributeKeyAccount, account),
		sdk.NewAttribute(AttributeKeyWId, wId),
		sdk.NewAttribute(AttributeKeyAsset, asset),
		sdk.NewAttribute(AttributeKeyQuantity, quantity),
		sdk.NewAttribute(AttributeKeyChainId, chainId),
	)
}
