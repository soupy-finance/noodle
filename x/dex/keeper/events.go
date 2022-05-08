package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/soupy-finance/noodle/x/dex/types"
)

func EmitAddOfferEvent(ctx sdk.Context, account sdk.AccAddress, market string, quantity, price sdk.Dec, side types.Side) {
	ctx.EventManager().EmitEvent(
		types.NewAddOfferEvent(account.String(), market, quantity, price, side),
	)
}

func EmitRemoveOfferEvent(ctx sdk.Context, account, market, quantity, price string, side types.Side) {
	ctx.EventManager().EmitEvent(
		types.NewRemoveOfferEvent(account, market, quantity, price, side),
	)
}

func EmitTradeExecEvent(ctx sdk.Context, maker, taker, market, quantity, price string) {
	ctx.EventManager().EmitEvent(
		types.NewTradeExecEvent(maker, taker, market, quantity, price),
	)
}

func EmitBalanceChangeEvent(ctx sdk.Context, account, asset, quantity string) {
	ctx.EventManager().EmitEvent(
		types.NewBalanceChangeEvent(account, asset, quantity),
	)
}
