package cmd

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/memoio/middleware-response/response"
	"github.com/urfave/cli/v2"
)

var signmCmd = &cli.Command{
	Name:      "signm",
	Usage:     "sign a message",
	ArgsUsage: "[sk][message]",
	Action: func(ctx *cli.Context) error {
		args := ctx.Args()
		if args.Len() != 2 {
			return fmt.Errorf("args number not right, need 4")
		}
		sk := args.Get(0)
		message := args.Get(1)

		skEcdsa, err := crypto.HexToECDSA(sk)
		if err != nil {
			return err
		}

		hash, err := hexutil.Decode(message)
		if err != nil {
			return err
		}
		sig, err := crypto.Sign(hash, skEcdsa)
		if err != nil {
			return err
		}
		sign := hexutil.Encode(sig)
		fmt.Println("sign message: ", sign)
		return nil
	},
}

var signTxCmd = &cli.Command{
	Name:      "signtx",
	Usage:     "sign a transaction",
	ArgsUsage: "[sk][transaction]",
	Action: func(ctx *cli.Context) error {
		args := ctx.Args()
		if args.Len() != 2 {
			return fmt.Errorf("args number not right, need 2")
		}
		sk := args.Get(0)
		message := args.Get(1)

		msgb, err := hexutil.Decode(message)
		if err != nil {
			return err
		}

		re := response.Transaction{}
		err = re.Unmarshal(msgb)
		if err != nil {
			return err
		}

		// txs, err := re.Marshal()
		// if err != nil {
		// 	return err
		// }
		fmt.Println("tx: ", string(msgb))

		tx := types.NewTransaction(re.Nonce, re.To, re.Value, re.Gas, re.GasPrice, re.Data)
		skEcdsa, err := crypto.HexToECDSA(sk)
		if err != nil {
			return err
		}

		signedTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(re.ChainId)), skEcdsa)
		if err != nil {
			return err
		}

		signb, err := signedTx.MarshalBinary()
		if err != nil {
			return err
		}

		signTx := response.SignTx{
			EndPoint: re.EndPoint,
			Tx:       signb,
		}
		signTxB, err := signTx.Marshal()
		if err != nil {
			return err
		}
		fmt.Println("sign transaction: ", hexutil.Encode(signTxB))
		return nil
	},
}
