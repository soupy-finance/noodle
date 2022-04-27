package keeper

import (
	"context"
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/soupy-finance/noodle/x/bridge/types"
)

func (k msgServer) ObserveDeposit(goCtx context.Context, msg *types.MsgObserveDeposit) (*types.MsgObserveDepositResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.stakingKeeper.GetValidator(ctx, sdk.ValAddress(msg.Creator))

	if !found || !val.IsBonded() {
		return nil, types.NotValidator
	}

	chain_contracts := k.ChainContracts(ctx)
	_, ok := chain_contracts[msg.ChainId]

	if !ok {
		return nil, types.InvalidChain
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.DepositsKey))
	depositorKeyBytes := []byte(msg.DepositId + ":" + msg.ChainId + ":" + msg.Depositor + ":" + msg.Quantity + ":" + msg.Asset)
	observationsBytes := store.Get(depositorKeyBytes)
	var observations []string

	if observationsBytes != nil {
		err := json.Unmarshal(observationsBytes, &observations)

		if err != nil {
			panic(err)
		}
	}

	observations = append(observations, msg.Creator)
	lastTotalPower := k.stakingKeeper.GetLastTotalPower(ctx)
	halfLastTotalPower := lastTotalPower.Quo(sdk.NewInt(2))
	var totalPower sdk.Int

	for _, validatorAddr := range observations {
		val, found := k.stakingKeeper.GetValidator(ctx, sdk.ValAddress(validatorAddr))

		if found {
			consensusPower := val.GetConsensusPower(val.GetBondedTokens())
			totalPower = totalPower.Add(sdk.NewInt(consensusPower))
		}
	}

	if totalPower.GT(halfLastTotalPower) {
		// Remove observations
		store.Delete(depositorKeyBytes)

		// Grant user tokens on exchange
		quantity, ok := sdk.NewIntFromString(msg.Quantity)

		if !ok {
			panic(types.InvalidQuantity)
		}

		coins := sdk.NewCoins(sdk.NewCoin(msg.Asset, quantity))
		err := k.bankKeeper.MintCoins(ctx, types.ModuleName, coins)

		if err != nil {
			panic(err)
		}

		err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sdk.AccAddress(msg.Depositor), coins)

		if err != nil {
			panic(err)
		}
	}

	return &types.MsgObserveDepositResponse{}, nil
}
