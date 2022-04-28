package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/soupy-finance/noodle/x/bridge/types"
)

func EmitWithdrawEvent(ctx sdk.Context, account sdk.AccAddress, asset string, quantity string, chainId string) {
	ctx.EventManager().EmitEvent(
		types.NewWithdrawEvent(account, asset, quantity, chainId),
	)
}
