package main

import "math/big"

// ResultV3 result for one pair
type ResultV3 struct {
	totalAmount0      *big.Float
	totalAmount1      *big.Float
	totalFee          *big.Int
	tokensOwed0       *big.Int
	Amounts           []LiqPositionAmount
	Token0            Token
	Token1            Token
	PoolAddr          string
	Slot0SqrtPriceX96 *big.Int
}

// ResultV2 result for one pair
type ResultV2 struct {
	totalAmount0      *big.Float
	totalAmount1      *big.Float
	totalLiqFromGraph *big.Float
	Amounts           []LiqPositionAmount
	Token0            Token
	Token1            Token
	PoolAddr          string
}

type LiqPositionAmount struct {
	Address             string `json:"address"`
	Token0              string `json:"token0"`
	Token1              string `json:"token1"`
	Token0HumanReadable string `json:"token0readable"`
	Token1HumanReadable string `json:"token1readable"`
}
