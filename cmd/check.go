package cmd

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/memoio/middleware-response/response"
	"github.com/urfave/cli/v2"
)

var checkCmd = &cli.Command{
	Name:      "check",
	Usage:     "get a check hash",
	ArgsUsage: "[pay][seller][size][nonce]",
	Action: func(ctx *cli.Context) error {
		args := ctx.Args()
		if args.Len() != 4 {
			return fmt.Errorf("args number not right, need 4")
		}
		a1 := args.Get(0)
		a2 := args.Get(1)
		a3 := args.Get(2)
		a4 := args.Get(3)

		pay := common.HexToAddress(a1)
		seller := common.HexToAddress(a2)
		size := toBigInt(a3).Uint64()
		nonce := toBigInt(a4)

		res := response.CheckResponse{
			PayAddr:  pay,
			Seller:   seller,
			SizeByte: size,
			Nonce:    nonce,
		}

		fmt.Println("check hash: ", hexutil.Encode(res.Hash()))
		return nil
	},
}
