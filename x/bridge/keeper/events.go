package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/soupy-finance/noodle/x/bridge/types"
)

func EmitWithdrawEvent(ctx sdk.Context, wId string, asset string, quantity string, address string, chainId string) {
	ctx.EventManager().EmitEvent(
		types.NewWithdrawEvent(wId, asset, quantity, address, chainId),
	)
}
