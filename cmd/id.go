package cmd

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/urfave/cli/v2"
)

var idCmd = &cli.Command{
	Name:      "id",
	Usage:     "get id info",
	ArgsUsage: "[sk][did][requestID]",
	Action: func(ctx *cli.Context) error {
		args := ctx.Args()
		if args.Len() != 3 {
			return fmt.Errorf("args number not right, need 3")
		}

		secretKey := args.Get(0)
		did := args.Get(1)
		requestID := args.Get(2)

		fmt.Println("sk:", secretKey)
		fmt.Println("did: ", did)
		fmt.Println("token: ", string(token))

		fmt.Println("requestID: ", requestID)
		id, err := strconv.ParseInt(requestID, 10, 64)
		if err != nil {
			return err
		}

		// create request with did, token ,sk
		fmt.Println("sign")
		err = sign(did, secretKey, id)
		if err != nil {
			return err
		}
		return nil
	},
}

func sign(did string, privateKeyHex string, requestID int64) error {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return err
	}

	// hash and sign
	hash := crypto.Keccak256([]byte(did), token, int64ToBytes(requestID))
	signature, err := crypto.Sign(hash, privateKey)
	if err != nil {
		return err
	}

	var payload = make(map[string]interface{})
	payload["did"] = did
	payload["token"] = hexutil.Encode(token)
	payload["requestID"] = requestID
	payload["signature"] = hexutil.Encode(signature)

	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	fmt.Println(string(b))

	return err
}
