package main

import (
	"math/big"
	"testing"
)

const (
	expectedTickAtSqrtPrice       = -1373112.1655059285
	expectedTokenAmount1          = 2407095255168192500
	expectedTokenAmount2          = 2.4070952551681923
	liquidityValue                = "12558033400096537032"
	sqrtPriceX96Value             = "2025953380162437579067355541581128"
	tokenPriceLow                 = 202980
	tokenPriceHigh                = 203040
	precision                     = 1000
	expectedTickAtSqrtPriceString = "-1373112.1655059285"
)

func TestGetTickAtSqrtPrice(t *testing.T) {
	res := getTickAtSqrtPrice(0.12123234234)
	if res != expectedTickAtSqrtPrice {
		t.Fatalf("Expected tick %f, got %f", expectedTickAtSqrtPrice, res)
	}

	input := new(big.Float).SetPrec(100).SetFloat64(0.12123234234)
	res2 := getTickAtSqrtPriceBigFloat(input)
	expectedTick := big.NewFloat(expectedTickAtSqrtPrice)
	if res2.Cmp(expectedTick) != 0 {
		t.Fatalf("Expected tick %s, got %s", expectedTick.String(), res2.String())
	}
}

func TestGetTokenAmounts(t *testing.T) {
	liquidity := newBigFloatOrFail(t, liquidityValue)
	sqrtPriceX96 := newBigFloatOrFail(t, sqrtPriceX96Value)

	am1, am2 := getTokenAmounts(liquidity, sqrtPriceX96, big.NewFloat(tokenPriceLow), big.NewFloat(tokenPriceHigh))
	checkAmount(t, am1, expectedTokenAmount1, "Token Amount 1")
	checkAmount(t, am2, expectedTokenAmount2, "Token Amount 2")
}

func TestGetTokenAmountsBig(t *testing.T) {
	liquidity := newBigFloatWithPrecisionOrFail(t, liquidityValue, precision)
	sqrtPriceX96 := newBigFloatWithPrecisionOrFail(t, sqrtPriceX96Value, precision)

	am1, am2 := getTokenAmountsBigFloat(liquidity, sqrtPriceX96,
		big.NewFloat(tokenPriceLow).SetPrec(precision),
		big.NewFloat(tokenPriceHigh).SetPrec(precision))
	checkAmount(t, am1, expectedTokenAmount1, "Token Amount 1")
	checkAmount(t, am2, expectedTokenAmount2, "Token Amount 2")
}

func newBigFloatOrFail(t *testing.T, value string) *big.Float {
	v, ok := new(big.Float).SetString(value)
	if !ok {
		t.Fatalf("Failed to create big.Float from string %s", value)
	}
	return v
}

func newBigFloatWithPrecisionOrFail(t *testing.T, value string, prec uint) *big.Float {
	v := newBigFloatOrFail(t, value)
	v.SetPrec(prec)
	return v
}

func checkAmount(t *testing.T, actual *big.Float, expected float64, label string) {
	expectedBig := big.NewFloat(expected)
	if actual.Cmp(expectedBig) != 0 {
		t.Fatalf("%s mismatch: expected %f, got %s", label, expected, actual.String())
	}
}
