package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/dex module sentinel errors
var (
	InvalidMarket       = sdkerrors.Register(ModuleName, 1100, "invalid market")
	InvalidOrderType    = sdkerrors.Register(ModuleName, 1101, "invalid order type")
	NoOrdersInBook      = sdkerrors.Register(ModuleName, 1102, "no orders in book")
	InsufficientBalance = sdkerrors.Register(ModuleName, 1103, "insufficient balance")
)
