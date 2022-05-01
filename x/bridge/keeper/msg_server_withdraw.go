package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/soupy-finance/noodle/x/bridge/types"
)

func (k msgServer) Withdraw(goCtx context.Context, msg *types.MsgWithdraw) (*types.MsgWithdrawResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check inputs
	chain_contracts := k.ChainContracts(ctx)
	_, ok := chain_contracts[msg.ChainId]

	if !ok {
		return nil, types.InvalidChain
	}

	quantity, ok := sdk.NewIntFromString(msg.Quantity)

	if !ok {
		return nil, types.InvalidQuantity
	}

	// Send coins to module and burn
	coins := sdk.NewCoins(sdk.NewCoin(msg.Asset, quantity))
	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, sdk.AccAddress(msg.Creator), types.ModuleName, coins)

	if err != nil {
		return nil, err
	}

	err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, coins)

	if err != nil {
		return nil, err
	}

	// Get withdrawal count
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.WithdrawalCountsKey))
	withdrawalCountKeyBytes := []byte(msg.Creator)
	withdrawalIdBytes := store.Get(withdrawalCountKeyBytes)

	if withdrawalIdBytes == nil {
		withdrawalIdBytes = []byte("0")
	}

	withdrawalId, ok := sdk.NewIntFromString(string(withdrawalIdBytes))

	if !ok {
		panic("invalid withdrawal count")
	}

	// Alert validators they need to sign
	EmitWithdrawEvent(ctx, msg.Creator, withdrawalId.String(), msg.Asset, msg.Quantity, msg.ChainId)
	// The withdrawal chain needs to store the withdrawal id once executed to prevent duplication

	// Increment and store withdrawal id
	withdrawalId.Add(sdk.NewInt(1))
	store.Set(withdrawalCountKeyBytes, []byte(withdrawalId.String()))

	return &types.MsgWithdrawResponse{}, nil
}
