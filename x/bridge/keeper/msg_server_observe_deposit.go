package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/soupy-finance/noodle/x/bridge/types"
)

func (k msgServer) ObserveDeposit(goCtx context.Context, msg *types.MsgObserveDeposit) (*types.MsgObserveDepositResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Parse inputs
	account, err := sdk.AccAddressFromBech32(msg.Creator)

	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	valAddr := sdk.ValAddress(account)
	val, found := k.stakingKeeper.GetValidator(ctx, valAddr)

	if !found || !val.IsBonded() {
		return nil, types.NotValidator
	}

	chain_contracts := k.ChainContractsParsed(ctx)
	_, ok := chain_contracts[msg.ChainId]

	if !ok {
		return nil, types.InvalidChain
	}

	depositor := k.DelegateAddressToDexAddress(ctx, msg.Depositor)
	depositorAddr, err := sdk.AccAddressFromBech32(depositor)

	if err != nil {
		return nil, types.InvalidDepositAddress
	}

	depositKeyBytes := GetDepositKeyBytes(depositor, msg)
	observations, err := k.GetDepositObservations(ctx, depositKeyBytes, msg)

	if err != nil {
		panic(err)
	}

	// Check if sufficient observation power reached
	observations = append(observations, valAddr.String())
	lastTotalPower := k.stakingKeeper.GetLastTotalPower(ctx)
	halfLastTotalPower := lastTotalPower.Quo(sdk.NewInt(2))
	totalPower := sdk.ZeroInt()
	validatorsObserved := make(map[string]bool, len(observations))

	for _, valAddrStr := range observations {
		if validatorsObserved[valAddrStr] {
			return nil, types.DupObservation
		}

		valAddr, err := sdk.ValAddressFromBech32(valAddrStr)

		if err != nil {
			panic(err)
		}

		val, found := k.stakingKeeper.GetValidator(ctx, sdk.ValAddress(valAddr))

		if found {
			validatorsObserved[valAddrStr] = true
			consensusPower := val.GetConsensusPower(k.stakingKeeper.PowerReduction(ctx))
			totalPower = totalPower.Add(sdk.NewInt(consensusPower))
		}
	}

	if totalPower.GT(halfLastTotalPower) {
		// Remove observations
		k.DeleteDepositObservations(ctx, depositKeyBytes, msg)

		// Grant user tokens on exchange
		quantity := sdk.MustNewDecFromStr(msg.Quantity)
		coins := sdk.NewCoins(sdk.NewCoin(msg.Asset, sdk.NewIntFromBigInt(quantity.BigInt())))
		err = k.bankKeeper.MintCoins(ctx, types.ModuleName, coins)

		if err != nil {
			panic(err)
		}

		err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, depositorAddr, coins)

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
