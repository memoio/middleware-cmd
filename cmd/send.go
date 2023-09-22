package cmd

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/memoio/middleware-response/response"
	"github.com/urfave/cli/v2"
)

var sendCmd = &cli.Command{
	Name:      "send",
	Usage:     "send transaction",
	ArgsUsage: "[transaction]",
	Action: func(ctx *cli.Context) error {
		args := ctx.Args()
		if args.Len() != 1 {
			return fmt.Errorf("args number not right, need 1")
		}

		signTxString := args.Get(0)
		signedTxBytes, err := hexutil.Decode(signTxString)
		if err != nil {
			return err
		}

		signTx := response.SignTx{}
		err = signTx.Unmarshal(signedTxBytes)
		if err != nil {
			return err
		}

		client, err := ethclient.DialContext(context.Background(), signTx.EndPoint)
		if err != nil {
			return err
		}

		defer client.Close()

		var signedTx = new(types.Transaction)
		err = signedTx.UnmarshalBinary(signTx.Tx)
		if err != nil {
			return err
		}

		err = client.SendTransaction(context.Background(), signedTx)
		if err != nil {
			return err
		}

		return nil
	},
}
