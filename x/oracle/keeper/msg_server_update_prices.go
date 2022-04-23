package keeper

import (
	"context"
	"encoding/binary"
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/soupy-finance/noodle/x/oracle/types"
)

func (k msgServer) UpdatePrices(goCtx context.Context, msg *types.MsgUpdatePrices) (*types.MsgUpdatePricesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Verify caller is a validator
	val, found := k.stakingKeeper.GetValidator(ctx, sdk.ValAddress(msg.Creator))

	if !found || !val.IsBonded() {
		return nil, types.NotValidator
	}

	// Unmarshal JSON
	var data map[string]uint64
	err := json.Unmarshal([]byte(msg.Data), &data)

	if err != nil {
		return nil, types.InvalidPriceData
	}

	// Iterate assets and update their prices for the validator
	assets := k.Assets(ctx)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.PricesKey))

	for _, asset := range assets {
		price, set := data[asset]

		if set {
			valAssetKeyBytes := []byte(msg.Creator + ":" + asset)
			priceBytes := make([]byte, 8)
			binary.BigEndian.PutUint64(priceBytes, price)
			store.Set(valAssetKeyBytes, priceBytes)
		}
	}

	return &types.MsgUpdatePricesResponse{}, nil
}
