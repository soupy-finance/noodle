package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/soupy-finance/noodle/x/oracle/types"
)

func EmitPricesEvents(ctx sdk.Context, prices string) {
	ctx.EventManager().EmitEvent(
		types.NewPricesEvent(prices),
	)
}
