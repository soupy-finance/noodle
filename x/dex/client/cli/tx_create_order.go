package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/soupy-finance/noodle/x/dex/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCreateOrder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-order [market] [side] [order-type] [price] [quantity]",
		Short: "Broadcast message createOrder",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argMarket := args[0]
			argSide := args[1]
			argOrderType := args[2]
			argPrice := args[3]
			argQuantity := args[4]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateOrder(
				clientCtx.GetFromAddress().String(),
				argMarket,
				argSide,
				argOrderType,
				argPrice,
				argQuantity,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
