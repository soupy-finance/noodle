package keeper_test

import (
	"testing"

	testkeeper "github.com/soupy-finance/noodle/testutil/keeper"
	"github.com/soupy-finance/noodle/x/oracle/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.OracleKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
	require.EqualValues(t, params.Assets, k.Assets(ctx))
}
