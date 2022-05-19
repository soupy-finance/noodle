package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/bridge module sentinel errors
var (
	NotValidator          = sdkerrors.Register(ModuleName, 1100, "not a validator")
	InvalidChain          = sdkerrors.Register(ModuleName, 1101, "invalid chain")
	InvalidQuantity       = sdkerrors.Register(ModuleName, 1102, "invalid quantity")
	InvalidChainAddress   = sdkerrors.Register(ModuleName, 1103, "invalid chain address")
	DupObservation        = sdkerrors.Register(ModuleName, 1104, "duplicate observation")
	InvalidDepositAddress = sdkerrors.Register(ModuleName, 1105, "invalid deposit address")
)
