package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"ledger/x/ledger/types"
)

var _ = strconv.Itoa(0)

func CmdCommitTstmt() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "commit-tstmt [payor] [payee] [denom] [q] [ts] [cid]",
		Short: "Broadcast message commit-tstmt",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argPayor := args[0]
			argPayee := args[1]
			argDenom := args[2]
			argQ, err := cast.ToUint64E(args[3])
			if err != nil {
				return err
			}
			argTs, err := cast.ToUint64E(args[4])
			if err != nil {
				return err
			}
			argCid := args[5]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCommitTstmt(
				clientCtx.GetFromAddress().String(),
				argPayor,
				argPayee,
				argDenom,
				argQ,
				argTs,
				argCid,
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
