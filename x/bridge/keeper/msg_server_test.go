package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/soupy-finance/noodle/testutil/keeper"
	"github.com/soupy-finance/noodle/x/bridge/keeper"
	"github.com/soupy-finance/noodle/x/bridge/types"
)

func setupMsgServer(
	t testing.TB,
	bankKeeper types.BankKeeper,
	stakingKeeper types.StakingKeeper,
) (keeper.Keeper, types.MsgServer, context.Context) {
	k, ctx := keepertest.BridgeKeeper(t, bankKeeper, stakingKeeper)
	return *k, keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
