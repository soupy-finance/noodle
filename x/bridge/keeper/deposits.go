package keeper

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/soupy-finance/noodle/x/bridge/types"
)

func (k Keeper) GetDepositObservations(
	ctx sdk.Context,
	depositKeyBytes []byte,
	msg *types.MsgObserveDeposit,
) (observations []string, err error) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.DepositsKey))
	observationsBytes := store.Get(depositKeyBytes)

	if observationsBytes != nil {
		err = json.Unmarshal(observationsBytes, &observations)
	}

	return
}

func (k Keeper) SetDepositObservations(
	ctx sdk.Context,
	depositKeyBytes []byte,
	msg *types.MsgObserveDeposit,
	observations []string,
) error {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.DepositsKey))
	observationsBytes, err := json.Marshal(observations)

	if err != nil {
		return err
	}

	store.Set(depositKeyBytes, observationsBytes)
	return nil
}

func (k Keeper) DeleteDepositObservations(ctx sdk.Context, depositKeyBytes []byte, msg *types.MsgObserveDeposit) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.DepositsKey))
	store.Delete(depositKeyBytes)
}

func (k Keeper) DelegateAddressToDexAddress(ctx sdk.Context, address string) string {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.AccountLinksKey))
	addressKeyBytes := []byte(address)
	dexAddressBytes := store.Get(addressKeyBytes)

	if dexAddressBytes == nil {
		return address
	}

	return string(dexAddressBytes)
}

func GetDepositKeyBytes(depositor string, msg *types.MsgObserveDeposit) []byte {
	return []byte(depositor + ":" + msg.DepositId + ":" + msg.ChainId + ":" + msg.Quantity + ":" + msg.Asset)
}
