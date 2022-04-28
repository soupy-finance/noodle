package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	EventTypeWithdraw = "withdraw"

	AttributeKeyAccount  = "account"
	AttributeKeyAsset    = "asset"
	AttributeKeyQuantity = "quantity"
	AttributeKeyChainId  = "chain_id"
)

func NewWithdrawEvent(account sdk.AccAddress, asset string, quantity string, chainId string) sdk.Event {
	return sdk.NewEvent(
		EventTypeWithdraw,
		sdk.NewAttribute(AttributeKeyAccount, account.String()),
		sdk.NewAttribute(AttributeKeyAsset, asset),
		sdk.NewAttribute(AttributeKeyQuantity, quantity),
		sdk.NewAttribute(AttributeKeyChainId, chainId),
	)
}
