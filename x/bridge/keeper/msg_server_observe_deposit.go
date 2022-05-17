package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/soupy-finance/noodle/x/bridge/types"
)

func (k msgServer) ObserveDeposit(goCtx context.Context, msg *types.MsgObserveDeposit) (*types.MsgObserveDepositResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.stakingKeeper.GetValidator(ctx, sdk.ValAddress(msg.Creator))

	if !found || !val.IsBonded() {
		return nil, types.NotValidator
	}

	chain_contracts := k.ChainContractsParsed(ctx)
	_, ok := chain_contracts[msg.ChainId]

	if !ok {
		return nil, types.InvalidChain
	}

	depositor := k.ExternalAddressToDexAddress(ctx, msg.Depositor)
	depositKeyBytes := GetDepositKeyBytes(depositor, msg)
	observations, err := k.GetDepositObservations(ctx, depositKeyBytes, msg)

	if err != nil {
		panic(err)
	}

	observations = append(observations, msg.Creator)
	lastTotalPower := k.stakingKeeper.GetLastTotalPower(ctx)
	halfLastTotalPower := lastTotalPower.Quo(sdk.NewInt(2))
	totalPower := sdk.ZeroInt()
	validatorsObserved := make(map[string]bool, len(observations))

	for _, validatorAddr := range observations {
		if validatorsObserved[validatorAddr] {
			return nil, types.DupObservation
		}

		val, found := k.stakingKeeper.GetValidator(ctx, sdk.ValAddress(validatorAddr))

		if found {
			validatorsObserved[validatorAddr] = true
			consensusPower := val.GetConsensusPower(val.GetBondedTokens())
			totalPower = totalPower.Add(sdk.NewInt(consensusPower))
		}
	}

	if totalPower.GT(halfLastTotalPower) {
		// Remove observations
		k.DeleteDepositObservations(ctx, depositKeyBytes, msg)

		// Grant user tokens on exchange
		quantity, err := sdk.NewDecFromStr(msg.Quantity)

		if err != nil {
			panic(err)
		}

		coins := sdk.NewCoins(sdk.NewCoin(msg.Asset, sdk.NewIntFromBigInt(quantity.BigInt())))
		err = k.bankKeeper.MintCoins(ctx, types.ModuleName, coins)

		if err != nil {
			panic(err)
		}

		err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sdk.AccAddress(depositor), coins)

		if err != nil {
			panic(err)
		}
	} else {
		err := k.SetDepositObservations(ctx, depositKeyBytes, msg, observations)

		if err != nil {
			panic(err)
		}
	}

	return &types.MsgObserveDepositResponse{}, nil
}
