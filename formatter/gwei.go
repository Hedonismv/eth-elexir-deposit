package formatter

import (
	"math/big"
	"math/rand"
	"time"
)

func GetRandomAmount(balance *big.Int, minPercent int, maxPercent int) *big.Int {
	rand.Seed(time.Now().UnixNano())
	percent := rand.Intn(maxPercent-minPercent) + minPercent
	percentFloat := float64(percent) / 100.0
	amountFloat := new(big.Float).Mul(new(big.Float).SetInt(balance), big.NewFloat(percentFloat))
	amount := new(big.Int)
	amountFloat.Int(amount)
	return amount
}

// ConvertWeiToEther  wei to ETH
func ConvertWeiToEther(weiAmount *big.Int) float64 {
	value := new(big.Float).SetInt(weiAmount)
	ethValue := new(big.Float).Quo(value, big.NewFloat(1e18))
	ethAmount, _ := ethValue.Float64()
	return ethAmount
}
