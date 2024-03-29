package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/oracle module sentinel errors
var (
	InvalidPriceData = sdkerrors.Register(ModuleName, 1100, "invalid price data")
	NotValidator     = sdkerrors.Register(ModuleName, 1101, "not a validator")
	InvalidAssetList = sdkerrors.Register(ModuleName, 1102, "invalid asset list")
	PriceAggError    = sdkerrors.Register(ModuleName, 1103, "price aggregation error")
)
