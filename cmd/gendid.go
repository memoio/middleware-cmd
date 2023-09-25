package cmd

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/memoio/go-did/memo"
	"github.com/urfave/cli/v2"
)

var gendidCmd = &cli.Command{
	Name:      "gendid",
	Usage:     "generate a did with an sk",
	ArgsUsage: "[sk]",
	Action: func(ctx *cli.Context) error {
		args := ctx.Args()
		if args.Len() != 1 {
			return errors.New("args number not right, need 1")
		}
		a1 := args.Get(0)

		bi, b := new(big.Int).SetString(a1, 16)
		if !b {
			return errors.New("set big int failed")
		}
		//fmt.Println("sk: ", bi.Text(16))

		sk, err := crypto.ToECDSA(bi.Bytes())
		if err != nil {
			return err
		}

		wallet := crypto.PubkeyToAddress(sk.PublicKey)
		fmt.Println("wallet address:", wallet.String())

		fmt.Println("create controller.")
		controller, err := memo.NewMemoDIDController(sk, "dev")
		if err != nil {
			return err
		}

		fmt.Println("register DID")
		err = controller.RegisterDID()
		if err != nil {
			return err
		}

		fmt.Println("DID: ")
		fmt.Println(controller.DID())
		return nil
	},
}
