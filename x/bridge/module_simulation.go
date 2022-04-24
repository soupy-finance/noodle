package bridge

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/soupy-finance/noodle/testutil/sample"
	bridgesimulation "github.com/soupy-finance/noodle/x/bridge/simulation"
	"github.com/soupy-finance/noodle/x/bridge/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = bridgesimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgObserveDeposit = "op_weight_msg_observe_deposit"
	// TODO: Determine the simulation weight value
	defaultWeightMsgObserveDeposit int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	bridgeGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&bridgeGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {
	bridgeParams := types.DefaultParams()
	return []simtypes.ParamChange{
		simulation.NewSimParamChange(types.ModuleName, string(types.KeyChainContracts), func(r *rand.Rand) string {
			return string(types.Amino.MustMarshalJSON(bridgeParams.ChainContracts))
		}),
	}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgObserveDeposit int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgObserveDeposit, &weightMsgObserveDeposit, nil,
		func(_ *rand.Rand) {
			weightMsgObserveDeposit = defaultWeightMsgObserveDeposit
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgObserveDeposit,
		bridgesimulation.SimulateMsgObserveDeposit(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
