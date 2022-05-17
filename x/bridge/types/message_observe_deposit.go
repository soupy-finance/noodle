package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgObserveDeposit = "observe_deposit"

var _ sdk.Msg = &MsgObserveDeposit{}

func NewMsgObserveDeposit(creator string, chainId string, depositor string, depositId string, quantity string, asset string) *MsgObserveDeposit {
	return &MsgObserveDeposit{
		Creator:   creator,
		ChainId:   chainId,
		Depositor: depositor,
		DepositId: depositId,
		Quantity:  quantity,
		Asset:     asset,
	}
}

func (msg *MsgObserveDeposit) Route() string {
	return RouterKey
}

func (msg *MsgObserveDeposit) Type() string {
	return TypeMsgObserveDeposit
}

func (msg *MsgObserveDeposit) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgObserveDeposit) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgObserveDeposit) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
