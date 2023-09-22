package cmd

import "math/big"

func toBigInt(s string) *big.Int {
	b := new(big.Int)
	b.SetString(s, 10)
	return b
}
