package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/soupy-finance/noodle/x/bridge/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdObserveDeposit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "observe-deposit [chain-id] [depositor]",
		Short: "Broadcast message observeDeposit",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argChainId := args[0]
			argDepositor := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgObserveDeposit(
				clientCtx.GetFromAddress().String(),
				argChainId,
				argDepositor,
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
