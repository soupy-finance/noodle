package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/soupy-finance/noodle/x/bridge/types"
)

func EmitWithdrawEvent(ctx sdk.Context, account string, wId string, asset string, quantity string, chainId string) {
	ctx.EventManager().EmitEvent(
		types.NewWithdrawEvent(account, wId, asset, quantity, chainId),
	)
}
