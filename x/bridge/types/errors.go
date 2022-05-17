package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/bridge module sentinel errors
var (
	NotValidator          = sdkerrors.Register(ModuleName, 1100, "not a validator")
	InvalidChain          = sdkerrors.Register(ModuleName, 1101, "invalid chain")
	DupObservation        = sdkerrors.Register(ModuleName, 1102, "duplicate observation")
	InvalidDepositAddress = sdkerrors.Register(ModuleName, 1103, "invalid deposit address")
)
