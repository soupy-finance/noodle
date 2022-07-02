package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/dex module sentinel errors
var (
	InvalidMarket         = sdkerrors.Register(ModuleName, 1100, "invalid market")
	InvalidOrderType      = sdkerrors.Register(ModuleName, 1101, "invalid order type")
	InvalidOrderFlags     = sdkerrors.Register(ModuleName, 1102, "invalid order flags")
	InsufficientBalance   = sdkerrors.Register(ModuleName, 1103, "insufficient balance for order")
	InsufficientLiquidity = sdkerrors.Register(ModuleName, 1104, "insufficient liquidity for market order")
	InvalidQueryAddress   = sdkerrors.Register(ModuleName, 1105, "invalid address queried")
)
