package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const TypeMsgUpdatePrices = "update_prices"

var _ sdk.Msg = &MsgUpdatePrices{}

func NewMsgUpdatePrices(creator string, data string) *MsgUpdatePrices {
	return &MsgUpdatePrices{
		Creator: creator,
		Data:    data,
	}
}

func (msg *MsgUpdatePrices) Route() string {
	return RouterKey
}

func (msg *MsgUpdatePrices) Type() string {
	return TypeMsgUpdatePrices
}

func (msg *MsgUpdatePrices) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdatePrices) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdatePrices) ValidateBasic() error {
	return nil
}
