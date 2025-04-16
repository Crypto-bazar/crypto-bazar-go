package utils

import "math/big"

func ToTokenFloat(amount *big.Int) float64 {
	amountFloat := new(big.Float).SetInt(amount)
	decimals := big.NewFloat(1e18)
	result, _ := new(big.Float).Quo(amountFloat, decimals).Float64()
	return result
}
