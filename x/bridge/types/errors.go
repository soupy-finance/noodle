package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/bridge module sentinel errors
var (
	NotValidator = sdkerrors.Register(ModuleName, 1100, "not a validator")
)