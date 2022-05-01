package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/dex module sentinel errors
var (
	InvalidSide   = sdkerrors.Register(ModuleName, 1100, "invalid book side")
	InvalidMarket = sdkerrors.Register(ModuleName, 1101, "invalid market")
)
