package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/urfave/cli/v2"
)

var token = []byte("memo.test")

var loginCmd = &cli.Command{
	Name:      "login",
	Usage:     "login with did",
	ArgsUsage: "[sk][did]",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "url",
			Aliases: []string{"u"},
			Usage:   "url of swag page",
			Value:   "http://103.39.231.220:18070/login",
		},
	},
	Action: func(ctx *cli.Context) error {
		args := ctx.Args()
		if args.Len() != 2 {
			return fmt.Errorf("args number not right, need 2")
		}

		url := ctx.String("url")

		//endpoint := args.Get(0)
		sk := args.Get(0)
		did := args.Get(1)

		fmt.Println("url:", url)
		fmt.Println("sk:", sk)
		fmt.Println("did: ", did)
		fmt.Println("token: ", string(token))

		// create request with did, token ,sk
		fmt.Println("generating request")
		req, err := GetLoginRequest(url, sk, did)
		if err != nil {
			return err
		}

		// send request or not
		b := true
		if b {
			// send request for response
			fmt.Println("sending request")
			client := &http.Client{Timeout: time.Minute}
			res, err := client.Do(req)
			if err != nil {
				return err
			}
			defer res.Body.Close()

			// reading from response
			fmt.Println("reading from response")
			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				return err
			}

			if res.StatusCode != http.StatusOK {
				log.Printf("respond code[%d]: %s\n", res.StatusCode, string(body))
			} else {
				log.Println(string(body))
			}
		}
		return nil
	},
}

func GetLoginRequest(url string, sk string, did string) (*http.Request, error) {
	var timestamp = time.Now().Unix()

	privateKey, err := crypto.HexToECDSA(sk)
	if err != nil {
		return nil, err
	}

	hash := crypto.Keccak256([]byte(did), token, int64ToBytes(timestamp))
	signature, err := crypto.Sign(hash, privateKey)
	if err != nil {
		return nil, err
	}

	var payload = make(map[string]interface{})
	payload["did"] = did
	payload["nonce"] = hexutil.Encode(token)
	payload["timestamp"] = timestamp
	payload["signature"] = hexutil.Encode(signature)

	b, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	fmt.Println("login info:")
	fmt.Println(string(b))

	req, err := http.NewRequest("POST", url, bytes.NewReader(b))

	return req, err
}

func int64ToBytes(v int64) []byte {
	return []byte{
		byte(0xff & v),
		byte(0xff & (v >> 8)),
		byte(0xff & (v >> 16)),
		byte(0xff & (v >> 24)),
		byte(0xff & (v >> 32)),
		byte(0xff & (v >> 40)),
		byte(0xff & (v >> 48)),
		byte(0xff & (v >> 56)),
	}
}
