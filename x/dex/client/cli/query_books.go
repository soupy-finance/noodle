package cli

import (
    "strconv"
	
	"github.com/spf13/cobra"
    "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/soupy-finance/noodle/x/dex/types"
)

var _ = strconv.Itoa(0)

func CmdBooks() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "books [market] [amm]",
		Short: "Query books",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			 reqMarket := args[0]
			 reqAmm := args[1]
			
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryBooksRequest{
				
                Market: reqMarket, 
                Amm: reqAmm, 
            }

            

			res, err := queryClient.Books(cmd.Context(), params)
            if err != nil {
                return err
            }

            return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

    return cmd
}