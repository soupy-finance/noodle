package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/soupy-finance/noodle/x/bridge/types"
)

func EmitWithdrawEvent(ctx sdk.Context, account sdk.AccAddress, wId string, asset string, quantity string, chainId string) {
	ctx.EventManager().EmitEvent(
		types.NewWithdrawEvent(account, wId, asset, quantity, chainId),
	)
}
