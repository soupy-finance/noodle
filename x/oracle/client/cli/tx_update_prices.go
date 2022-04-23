package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/soupy-finance/noodle/x/oracle/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdUpdatePrices() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-prices [data]",
		Short: "Broadcast message updatePrices",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argData := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdatePrices(
				clientCtx.GetFromAddress().String(),
				argData,
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
