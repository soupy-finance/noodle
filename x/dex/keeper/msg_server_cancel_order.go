package keeper

import (
	"context"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/soupy-finance/noodle/x/dex/types"
)

func (k msgServer) CancelOrder(goCtx context.Context, msg *types.MsgCancelOrder) (*types.MsgCancelOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	account, err := sdk.AccAddressFromBech32(msg.Creator)

	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	var side types.Side

	if !msg.Side {
		side = 'b'
	} else {
		side = 'a'
	}

	markets := k.MarketsParsed(ctx)
	_, ok := markets[msg.Market]

	if !ok {
		return nil, types.InvalidMarket
	}

	price, err := sdk.NewDecFromStr(msg.Price)

	if err != nil {
		return nil, err
	}

	order := types.Order{
		Id:      types.OrderId(msg.Id),
		Account: account,
		Market:  msg.Market,
		Side:    side,
		Price:   price,
	}

	book, err := k.GetPureBook(ctx, order.Market, order.Side)

	if err != nil {
		return nil, err
	}

	assets := strings.Split(order.Market, "-")

	for i := range book.Levels {
		level := &book.Levels[i]

		if level.Price.Equal(order.Price) {
			for j := range level.Orders {
				_order := &level.Orders[j]

				if _order.Id == order.Id {
					level.Orders = append(level.Orders[:j], level.Orders[j+1:]...)
					var coins sdk.Coins

					if side == 'b' {
						coins = sdk.NewCoins(sdk.NewCoin(assets[1], sdk.NewIntFromBigInt(_order.Quantity.Mul(_order.Price).BigInt())))
					} else {
						coins = sdk.NewCoins(sdk.NewCoin(assets[0], sdk.NewIntFromBigInt(_order.Quantity.BigInt())))
					}

					k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, order.Account, coins)
					EmitRemoveOfferEvent(ctx, _order.Id, _order.Account, order.Market)
					break
				}
			}
		}
	}

	err = k.SaveVirtualBook(ctx, book)

	if err != nil {
		return nil, err
	}

	err = k.RemoveAccountOrder(ctx, &order)

	if err != nil {
		return nil, err
	}

	return &types.MsgCancelOrderResponse{}, nil
}
