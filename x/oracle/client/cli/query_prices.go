package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/soupy-finance/noodle/x/oracle/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdPrices() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "prices [assets]",
		Short: "Query prices",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqAssets := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryPricesRequest{

				Assets: reqAssets,
			}

			res, err := queryClient.Prices(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
