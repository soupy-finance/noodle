package oracle_test

import (
	"testing"

	keepertest "github.com/soupy-finance/noodle/testutil/keeper"
	"github.com/soupy-finance/noodle/testutil/nullify"
	"github.com/soupy-finance/noodle/x/oracle"
	"github.com/soupy-finance/noodle/x/oracle/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.OracleKeeper(t)
	oracle.InitGenesis(ctx, *k, genesisState)
	got := oracle.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
