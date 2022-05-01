package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/dex module sentinel errors
var (
	InvalidMarket    = sdkerrors.Register(ModuleName, 1100, "invalid market")
	InvalidPrice     = sdkerrors.Register(ModuleName, 1101, "invalid price")
	InvalidQuantity  = sdkerrors.Register(ModuleName, 1102, "invalid quantity")
	InvalidOrderType = sdkerrors.Register(ModuleName, 1103, "invalid order type")
)
