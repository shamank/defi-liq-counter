package main

import (
	"github.com/ALTree/bigfloat"
	"math"
	"math/big"
)

// https://blog.uniswap.org/uniswap-v3-math-primer-2

var Q96 = bigfloat.Pow(big.NewFloat(2), big.NewFloat(96)).SetPrec(prec)     // const Q96
var Q96float, _ = bigfloat.Pow(big.NewFloat(2), big.NewFloat(96)).Float64() // const Q96
var SpecialConst = big.NewFloat(1.0001).SetPrec(prec)

func getTickAtSqrtPrice(sqrtPriceX96 float64) float64 {
	return math.Log(math.Pow(sqrtPriceX96/Q96float, 2)) / math.Log(1.0001)
}

func getTickAtSqrtPriceBigFloat(sqrtPriceX96 *big.Float) *big.Float {
	ratio := new(big.Float).Quo(sqrtPriceX96, Q96).SetPrec(prec)   // sqrtPriceX96 / Q96
	ratioSquared := new(big.Float).Mul(ratio, ratio).SetPrec(prec) // (sqrtPriceX96 / Q96)^2

	logRatioSquared := bigfloat.Log(ratioSquared).SetPrec(prec)

	logSpecialConst := bigfloat.Log(SpecialConst).SetPrec(prec)

	tick := new(big.Float).Quo(logRatioSquared, logSpecialConst).SetPrec(prec)

	return tick
}

// (liquidity * ((sqrtRatioB - sqrtRatioA) / (sqrtRatioA * sqrtRatioB)))
func calcAmount0(liquidity, sqrtRatioB, sqrtRatioA *big.Float) *big.Float {
	temp1 := new(big.Float).Sub(sqrtRatioB, sqrtRatioA)
	temp2 := new(big.Float).Mul(sqrtRatioA, sqrtRatioB)
	result := new(big.Float).Quo(temp1, temp2)
	return result.Mul(result, liquidity)
}

// (liquidity * (sqrtRatioB - sqrtRatioA))
func calcAmount1(liquidity, sqrtRatioB, sqrtRatioA *big.Float) *big.Float {
	temp1 := new(big.Float).Sub(sqrtRatioB, sqrtRatioA)
	return new(big.Float).Mul(liquidity, temp1)
}

func getTokenAmounts(liquidity, sqrtPriceX96, tickLow, tickHigh *big.Float) (*big.Float, *big.Float) {

	// Pow uses the first argument's precision.
	up := bigfloat.Pow(SpecialConst, tickHigh)               // 1.0001^tickUpper
	low := bigfloat.Pow(SpecialConst, tickLow).SetPrec(prec) // 1.0001^tickLow
	sqrtRatioB := up.Sqrt(up)
	sqrtRatioA := low.Sqrt(low)

	sqrtPriceX96Float, _ := sqrtPriceX96.Float64()
	currentTick := big.NewFloat(getTickAtSqrtPrice(sqrtPriceX96Float))

	var sqrtPrice = sqrtPriceX96.Quo(sqrtPriceX96, Q96) // div by 2^96

	var amount0 = big.NewFloat(0)
	var amount1 = big.NewFloat(0)
	if currentTick.Cmp(tickLow) == -1 {
		amount0 = calcAmount0(liquidity, sqrtRatioB, sqrtRatioA)
	} else if currentTick.Cmp(tickHigh) >= 0 {
		amount1 = calcAmount1(liquidity, sqrtRatioB, sqrtRatioA)
	} else if currentTick.Cmp(tickLow) >= 0 && currentTick.Cmp(tickHigh) == -1 {
		amount0 = calcAmount0(liquidity, sqrtRatioB, sqrtPrice)
		amount1 = calcAmount1(liquidity, sqrtPrice, sqrtRatioA)
	}
	return amount0, amount1
}

func getTokenAmountsBigFloat(liquidity, sqrtPriceX96, tickLow, tickHigh *big.Float) (*big.Float, *big.Float) {

	// Pow uses the first argument's precision.
	up := bigfloat.Pow(SpecialConst, tickHigh)               // 1.0001^tickUpper
	low := bigfloat.Pow(SpecialConst, tickLow).SetPrec(prec) // 1.0001^tickLow
	sqrtRatioB := up.Sqrt(up)
	sqrtRatioA := low.Sqrt(low)

	//sqrtPriceX96Float, _ := sqrtPriceX96.Float64()
	currentTick := getTickAtSqrtPriceBigFloat(sqrtPriceX96)

	var sqrtPrice = sqrtPriceX96.Quo(sqrtPriceX96, Q96) // div by 2^96

	var amount0 = big.NewFloat(0)
	var amount1 = big.NewFloat(0)
	if currentTick.Cmp(tickLow) == -1 {
		amount0 = calcAmount0(liquidity, sqrtRatioB, sqrtRatioA)
	} else if currentTick.Cmp(tickHigh) >= 0 {
		amount1 = calcAmount1(liquidity, sqrtRatioB, sqrtRatioA)
	} else if currentTick.Cmp(tickLow) >= 0 && currentTick.Cmp(tickHigh) == -1 {
		amount0 = calcAmount0(liquidity, sqrtRatioB, sqrtPrice)
		amount1 = calcAmount1(liquidity, sqrtPrice, sqrtRatioA)
	}
	return amount0, amount1
}
