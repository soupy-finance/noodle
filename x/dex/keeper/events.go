package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/soupy-finance/noodle/x/dex/types"
)

func EmitAddOfferEvent(ctx sdk.Context, account sdk.AccAddress, id, market string, quantity, price sdk.Dec, side types.Side) {
	ctx.EventManager().EmitEvent(
		types.NewAddOfferEvent(account, id, market, quantity, price, side),
	)
}

func EmitRemoveOfferEvent(ctx sdk.Context, account sdk.AccAddress, id string) {
	ctx.EventManager().EmitEvent(
		types.NewRemoveOfferEvent(account, id),
	)
}

func EmitUpdateOfferEvent(ctx sdk.Context, account sdk.AccAddress, id string, quantity sdk.Dec) {
	ctx.EventManager().EmitEvent(
		types.NewUpdateOfferEvent(account, id, quantity),
	)
}

func EmitTradeExecEvent(ctx sdk.Context, maker, taker sdk.AccAddress, market string, quantity, price sdk.Dec) {
	ctx.EventManager().EmitEvent(
		types.NewTradeExecEvent(maker, taker, market, quantity, price),
	)
}

func EmitBalanceChangeEvent(ctx sdk.Context, account sdk.AccAddress, asset string, quantity sdk.Dec) {
	ctx.EventManager().EmitEvent(
		types.NewBalanceChangeEvent(account, asset, quantity),
	)
}
