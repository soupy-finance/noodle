package keeper

import (
	"context"
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/soupy-finance/noodle/x/oracle/types"
)

func (k msgServer) UpdatePrices(goCtx context.Context, msg *types.MsgUpdatePrices) (*types.MsgUpdatePricesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Parse inputs
	accAddr, err := sdk.AccAddressFromBech32(msg.Creator)

	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	// Verify caller is a validator
	valAddr := sdk.ValAddress(accAddr)
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

	for _, asset := range assets {
		priceStr, exists := data[asset]

		if !exists {
			continue
		}

		price, err := sdk.NewDecFromStr(priceStr)

		if err != nil {
			continue
		}

		k.valPrices[valAddr.String()+":"+asset] = price
	}

	return &types.MsgUpdatePricesResponse{}, nil
}
