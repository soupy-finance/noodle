package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/soupy-finance/noodle/x/dex/types"
)

func EmitAddOfferEvent(ctx sdk.Context, id types.OrderId, account sdk.AccAddress, market string, quantity, price sdk.Dec, side types.Side) {
	ctx.EventManager().EmitEvent(
		types.NewAddOfferEvent(id, account, market, quantity, price, side),
	)
}

func EmitRemoveOfferEvent(ctx sdk.Context, id types.OrderId, account sdk.AccAddress, market string) {
	ctx.EventManager().EmitEvent(
		types.NewRemoveOfferEvent(id, account, market),
	)
}

func EmitUpdateOfferEvent(ctx sdk.Context, id types.OrderId, account sdk.AccAddress, market string, quantity sdk.Dec) {
	ctx.EventManager().EmitEvent(
		types.NewUpdateOfferEvent(id, account, market, quantity),
	)
}

func EmitTradeExecEvent(ctx sdk.Context, maker, taker sdk.AccAddress, market string, quantity, price sdk.Dec, side types.Side) {
	ctx.EventManager().EmitEvent(
		types.NewTradeExecEvent(maker, taker, market, quantity, price, side),
	)
}

func EmitBalanceChangeEvent(ctx sdk.Context, account sdk.AccAddress, asset string, quantity sdk.Dec) {
	ctx.EventManager().EmitEvent(
		types.NewBalanceChangeEvent(account, asset, quantity),
	)
}
