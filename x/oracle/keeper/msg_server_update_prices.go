package keeper

import (
	"context"
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/soupy-finance/noodle/x/oracle/types"
)

func (k msgServer) UpdatePrices(goCtx context.Context, msg *types.MsgUpdatePrices) (*types.MsgUpdatePricesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Parse inputs
	valAddr, err := sdk.ValAddressFromBech32(msg.Creator)

	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	// Verify caller is a validator
	val, found := k.stakingKeeper.GetValidator(ctx, valAddr)

	if !found || !val.IsBonded() {
		return nil, types.NotValidator
	}

	// Unmarshal JSON
	var data map[string]string
	err = json.Unmarshal([]byte(msg.Data), &data)

	if err != nil {
		return nil, types.InvalidPriceData
	}

	// Iterate assets and update their prices for the validator
	assets := k.AssetsParsed(ctx)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.PricesKey))

	for _, asset := range assets {
		price, set := data[asset]

		if set {
			valAssetKeyBytes := append(valAddr, []byte(asset)...)
			store.Set(valAssetKeyBytes, []byte(price))
		}
	}

	return &types.MsgUpdatePricesResponse{}, nil
}
