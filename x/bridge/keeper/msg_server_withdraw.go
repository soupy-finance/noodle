package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/soupy-finance/noodle/x/bridge/types"
)

func (k msgServer) Withdraw(goCtx context.Context, msg *types.MsgWithdraw) (*types.MsgWithdrawResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Parse inputs
	chain_contracts := k.ChainContractsParsed(ctx)
	_, ok := chain_contracts[msg.ChainId]

	if !ok {
		return nil, types.InvalidChain
	}

	quantity, err := sdk.NewDecFromStr(msg.Quantity)

	if err != nil {
		return nil, types.InvalidQuantity
	}

	dexAddr, err := sdk.AccAddressFromBech32(msg.Creator)

	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	// Send coins to module and burn
	coins := sdk.NewCoins(sdk.NewCoin(msg.Asset, sdk.NewIntFromBigInt(quantity.BigInt())))
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, dexAddr, types.ModuleName, coins)

	if err != nil {
		return nil, err
	}

	err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, coins)

	if err != nil {
		return nil, err
	}

	// Get withdrawal count
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.WithdrawalCountsKey))
	withdrawalIdBytes := store.Get(dexAddr)

	if withdrawalIdBytes == nil {
		withdrawalIdBytes = []byte("0")
	}

	withdrawalId, ok := sdk.NewIntFromString(string(withdrawalIdBytes))

	if !ok {
		panic("invalid withdrawal count")
	}

	// Alert validators they need to sign
	EmitWithdrawEvent(ctx, withdrawalId.String(), msg.Asset, msg.Quantity, msg.Address, msg.ChainId)
	// The withdrawal chain needs to store the withdrawal id once executed to prevent duplication

	// Increment and store withdrawal id
	withdrawalId.Add(sdk.NewInt(1))
	store.Set(dexAddr, []byte(withdrawalId.String()))

	return &types.MsgWithdrawResponse{}, nil
}
