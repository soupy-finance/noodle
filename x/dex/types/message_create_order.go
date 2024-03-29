package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const TypeMsgCreateOrder = "create_order"

var _ sdk.Msg = &MsgCreateOrder{}

func NewMsgCreateOrder(creator string, market string, side bool, orderType string, price string, quantity string) *MsgCreateOrder {
	return &MsgCreateOrder{
		Creator:   creator,
		Market:    market,
		Side:      side,
		OrderType: orderType,
		Price:     price,
		Quantity:  quantity,
	}
}

func (msg *MsgCreateOrder) Route() string {
	return RouterKey
}

func (msg *MsgCreateOrder) Type() string {
	return TypeMsgCreateOrder
}

func (msg *MsgCreateOrder) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateOrder) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateOrder) ValidateBasic() error {
	return nil
}
