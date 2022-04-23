package oracle

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/soupy-finance/noodle/testutil/sample"
	oraclesimulation "github.com/soupy-finance/noodle/x/oracle/simulation"
	"github.com/soupy-finance/noodle/x/oracle/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = oraclesimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgUpdatePrices = "op_weight_msg_update_prices"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdatePrices int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	oracleGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&oracleGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {
	oracleParams := types.DefaultParams()
	return []simtypes.ParamChange{
		simulation.NewSimParamChange(types.ModuleName, string(types.KeyAssets), func(r *rand.Rand) string {
			return string(types.Amino.MustMarshalJSON(oracleParams.Assets))
		}),
	}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgUpdatePrices int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdatePrices, &weightMsgUpdatePrices, nil,
		func(_ *rand.Rand) {
			weightMsgUpdatePrices = defaultWeightMsgUpdatePrices
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdatePrices,
		oraclesimulation.SimulateMsgUpdatePrices(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
