package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	testkeeper "github.com/soupy-finance/noodle/testutil/keeper"
	"github.com/soupy-finance/noodle/x/oracle/types"
	"github.com/stretchr/testify/require"
)

func TestPricesQuery(t *testing.T) {
	keeper, ctx := testkeeper.OracleKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	params := types.DefaultParams()
	keeper.SetParams(ctx, params)

	request := types.QueryPricesRequest{
		Assets: []string{"eth"},
	}
	response, err := keeper.Prices(wctx, &request)
	require.NoError(t, err)
	require.Equal(t, "[\"0\"]", response.Data)
}
