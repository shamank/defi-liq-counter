// pancake3 bsc
//go:generate abigen --abi pancake3/bsc/nonfungiblePositionManager/abi.json --pkg nonfungiblePositionManager --bin pancake3/bsc/nonfungiblePositionManager/bytecode.bin --out pancake3/bsc/nonfungiblePositionManager/nonfungiblePositionManager.go
//go:generate abigen --abi pancake3/bsc/pancakeV3factory/abi.json --pkg pancakeV3factory --bin pancake3/bsc/pancakeV3factory/bytecode.bin --out pancake3/bsc/pancakeV3factory/pancakeV3factory.go
//go:generate abigen --abi pancake3/bsc/pancakeV3pool/abi.json --pkg pancakeV3pool --bin pancake3/bsc/pancakeV3pool/bytecode.bin --out pancake3/bsc/pancakeV3pool/pancakeV3pool.go

// uniswap3 bsc
//go:generate abigen --abi uniswap3/bsc/nonfungiblePositionManager/abi.json --pkg nonfungiblePositionManager --bin uniswap3/bsc/nonfungiblePositionManager/bytecode.bin --out uniswap3/bsc/nonfungiblePositionManager/nonfungiblePositionManager.go
//go:generate abigen --abi uniswap3/bsc/uniswapV3factory/abi.json --pkg uniswapV3factory --bin uniswap3/bsc/uniswapV3factory/bytecode.bin --out uniswap3/bsc/uniswapV3factory/uniswapV3factory.go
//go:generate abigen --abi uniswap3/bsc/uniswapV3pool/abi.json --pkg uniswapV3pool --bin uniswap3/bsc/uniswapV3pool/bytecode.bin --out uniswap3/bsc/uniswapV3pool/uniswapV3pool.go

// uniswap3 eth
//go:generate abigen --abi uniswap3/eth/nonfungiblePositionManager/abi.json --pkg nonfungiblePositionManager --bin uniswap3/eth/nonfungiblePositionManager/bytecode.bin --out uniswap3/eth/nonfungiblePositionManager/nonfungiblePositionManager.go
//go:generate abigen --abi uniswap3/eth/uniswapV3factory/abi.json --pkg uniswapV3factory --bin uniswap3/eth/uniswapV3factory/bytecode.bin --out uniswap3/eth/uniswapV3factory/uniswapV3factory.go
//go:generate abigen --abi uniswap3/eth/uniswapV3pool/abi.json --pkg uniswapV3pool --bin uniswap3/eth/uniswapV3pool/bytecode.bin --out uniswap3/eth/uniswapV3pool/uniswapV3pool.go

// uniswap2 eth
//go:generate abigen --abi uniswap2/eth/uniswapV2pair/abi.json --pkg uniswapV2pair --bin uniswap2/eth/uniswapV2pair/bytecode.bin --out uniswap2/eth/uniswapV2pair/uniswapV2pair.go

// pancake2 bsc
//go:generate abigen --abi pancake2/bsc/pancakeV2pair/abi.json --pkg pancakeV2pair --bin pancake2/bsc/pancakeV2pair/bytecode.bin --out pancake2/bsc/pancakeV2pair/pancakeV2pair.go

package main

import (
	npmPancake3Bsc "defi-liq-counter/pancake3/bsc/NonfungiblePositionManager"
	npmUniswap3Bsc "defi-liq-counter/uniswap3/bsc/NonfungiblePositionManager"
	npmUniswap3Eth "defi-liq-counter/uniswap3/eth/NonfungiblePositionManager"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math"
	"math/big"
	"os"
	"time"
)

const prec = 1000 // in binary digits

type App struct {
	Config
	ethClient  *ethclient.Client // evm
	opts       *bind.CallOpts
	posManager posManager
}

func main() {
	app := newAppFromConf()
	app.InitEVMclient()

	fmt.Println("DEX App:", app.DexApp)

	// Установка блок-опций, если указан номер блока
	if app.BlockNumber != 0 {
		blockNumber := big.NewInt(int64(app.BlockNumber))
		fmt.Printf("Using block %d\n", blockNumber)
		app.opts = &bind.CallOpts{BlockNumber: blockNumber}
	}

	owners := make(map[string]any)
	url := "https://etherscan.io"

	switch app.DexApp {
	case Pancake2bsc:
		url = "https://bscscan.com"
		processV2(app, owners, url)
	case Uniswap2eth:
		processV2(app, owners, url)
	case Uniswap3eth, Pancake3bsc, Uniswap3bsc:
		processV3(app, owners)
	default:
		log.Fatalf("Invalid DEX app: %s", app.DexApp)
	}
}

func processV2(app *App, owners map[string]any, url string) {
	for _, pair := range app.Pairs {
		// Получение владельцев через скан
		app.GetOwnersFromScanV2(owners, url, pair.PoolAddress)
		if len(owners) == 0 {
			fmt.Printf("%s: No positions (addresses) for pair %s-%s\n", url[8:], pair.Token0.Symbol, pair.Token1.Symbol)
		}

		// Получение владельцев через GraphQL
		app.GetOwnersFromGraphV2(owners, pair.Token0.Address, pair.Token1.Address)
		if len(owners) == 0 {
			fmt.Printf("No positions for pair %s-%s from Graph\n", pair.Token0.Symbol, pair.Token1.Symbol)
		}
	}

	fmt.Printf("TOTAL: %d unique addresses for all pairs\n", len(owners))
	resV2 := app.CountV2(owners)
	for _, pairResult := range resV2 {
		app.SaveV2Result(pairResult)
	}
}

func processV3(app *App, owners map[string]any) {
	app.InitPosManagerV3()

	for _, pair := range app.Pairs {
		// Получение владельцев через V3
		app.GetOwnersV3(owners, pair.Token0.Symbol, pair.Token1.Symbol)
		if len(owners) == 0 {
			fmt.Printf("No positions for pair %s-%s\n", pair.Token0.Symbol, pair.Token1.Symbol)
		}

		// Получение снапшотов владельцев
		app.GetOwnersV3Snapshots(owners, pair.PoolAddress)
		if len(owners) == 0 {
			fmt.Printf("No positions for pool %s\n", pair.PoolAddress)
		}
	}

	fmt.Printf("TOTAL: %d unique addresses for all pairs\n", len(owners))
	v3 := app.InitV3(owners)
	app.StartSyncV3(v3)
	app.CountV3(v3)

	for _, pairResult := range v3.result {
		app.SaveV3Result(pairResult)
	}
}

func (app *App) InitEVMclient() {
	var err error
	app.ethClient, err = ethclient.Dial(app.BlockchainURL)
	if err != nil {
		log.Fatalln("Can't init eth client: ", err)
	}
}

type posManager interface {
	BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error)
	TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error)
	Positions(opts *bind.CallOpts, tokenId *big.Int) (struct {
		Nonce                    *big.Int
		Operator                 common.Address
		Token0                   common.Address
		Token1                   common.Address
		Fee                      *big.Int
		TickLower                *big.Int
		TickUpper                *big.Int
		Liquidity                *big.Int
		FeeGrowthInside0LastX128 *big.Int
		FeeGrowthInside1LastX128 *big.Int
		TokensOwed0              *big.Int
		TokensOwed1              *big.Int
	}, error)
}

func (app *App) InitPosManagerV3() {
	var err error
	var uniswapV3npmAddrETH = common.HexToAddress("0xC36442b4a4522E871399CD717aBDD847Ab11FE88")
	var uniswapV3npmAddrBSC = common.HexToAddress("0x7b8A01B39D58278b5DE7e48c8449c9f4F5170613")
	var pancakev3npmAddr = common.HexToAddress("0x46A15B0b27311cedF172AB29E4f4766fbE7F4364")

	switch app.DexApp {
	case Pancake3bsc:
		app.posManager, err = npmPancake3Bsc.NewNonfungiblePositionManager(pancakev3npmAddr, app.ethClient)
		if err != nil {
			log.Fatalln("can't init position manager contract: ", err)
		}
	case Uniswap3bsc:
		app.posManager, err = npmUniswap3Bsc.NewNonfungiblePositionManager(uniswapV3npmAddrBSC, app.ethClient)
		if err != nil {
			log.Fatalln("can't init position manager contract: ", err)
		}
	case Uniswap3eth:
		app.posManager, err = npmUniswap3Eth.NewNonfungiblePositionManager(uniswapV3npmAddrETH, app.ethClient)
		if err != nil {
			log.Fatalln("can't init position manager contract: ", err)
		}
	case Uniswap2eth:
		// no need for this dex
	case Pancake2bsc:
		// no need for this dex
	default:
		log.Fatalln("invalid dex app network")
	}
}

func (app *App) SaveV2Result(pairResult *ResultV2) {
	filename := fmt.Sprintf("%s-%s-%s_%s_%s.json",
		pairResult.Token0.Symbol, pairResult.Token1.Symbol, pairResult.PoolAddr[38:],
		app.DexApp.String(),
		time.Now().Format("2006-01-02"))

	marshal, err := json.Marshal(pairResult.Amounts)
	if err != nil {
		log.Println("Can't json marshal result: ", err)
	}
	err = os.WriteFile(filename, marshal, 0666)
	if err != nil {
		log.Println("Can't write result to file: ", err)
	}
	fmt.Printf("-------- %s-%s --------\n", pairResult.Token0.Symbol, pairResult.Token1.Symbol)
	fmt.Println("poolAddr: ", pairResult.PoolAddr)
	fmt.Printf("totalAmount0 %s: %f\n", pairResult.Token0.Symbol, pairResult.totalAmount0)
	fmt.Printf("totalAmount1 %s: %f\n", pairResult.Token1.Symbol, pairResult.totalAmount1)

	// convert to human-readable
	fmt.Printf("totalAmount0 / 10^%d: %f %s\n", pairResult.Token0.Decimals, humanReadableFloat(pairResult.totalAmount0, pairResult.Token0.Decimals), pairResult.Token0.Symbol)
	fmt.Printf("totalAmount1 / 10^%d: %f %s\n", pairResult.Token1.Decimals, humanReadableFloat(pairResult.totalAmount1, pairResult.Token1.Decimals), pairResult.Token1.Symbol)
	fmt.Println("totalLiqFromGraph (from graph):", pairResult.totalLiqFromGraph)
}

func (app *App) SaveV3Result(pairResult *ResultV3) {
	filename := generateFilename(pairResult, app.DexApp)
	if err := saveResultToFile(pairResult.Amounts, filename); err != nil {
		log.Println("Error saving result to file:", err)
		return
	}

	printV3Result(pairResult)
}

func generateFilename(pairResult *ResultV3, dexApp DexApp) string {
	return fmt.Sprintf("%s-%s-%s_%s_%s.json",
		pairResult.Token0.Symbol,
		pairResult.Token1.Symbol,
		pairResult.PoolAddr[38:], // last 38 characters of PoolAddr
		dexApp.String(),
		time.Now().Format("2006-01-02"))
}

func saveResultToFile(data interface{}, filename string) error {
	marshal, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}
	if err := os.WriteFile(filename, marshal, 0666); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}
	return nil
}

func printV3Result(pairResult *ResultV3) {
	fmt.Printf("-------- %s-%s --------\n", pairResult.Token0.Symbol, pairResult.Token1.Symbol)
	fmt.Println("Pool Address:", pairResult.PoolAddr)

	fmt.Printf("Total Amount0 (%s): %f\n", pairResult.Token0.Symbol, pairResult.totalAmount0)
	fmt.Printf("Total Amount1 (%s): %f\n", pairResult.Token1.Symbol, pairResult.totalAmount1)

	// Convert to human-readable format
	totalAmount0HR := humanReadableFloat(pairResult.totalAmount0, pairResult.Token0.Decimals)
	totalAmount1HR := humanReadableFloat(pairResult.totalAmount1, pairResult.Token1.Decimals)

	fmt.Printf("Total Amount0 / 10^%d: %f %s\n", pairResult.Token0.Decimals, totalAmount0HR, pairResult.Token0.Symbol)
	fmt.Printf("Total Amount1 / 10^%d: %f %s\n", pairResult.Token1.Decimals, totalAmount1HR, pairResult.Token1.Symbol)

	fmt.Println("Total Tokens Owed0:", pairResult.tokensOwed0)
	fmt.Println("Total Fee:", pairResult.totalFee)
}

func humanReadableInt(amount *big.Int, decimals uint) *big.Int {
	return amount.Quo(amount, big.NewInt(int64(math.Pow(10, float64(decimals)))))
}

func humanReadableFloat(amount *big.Float, decimals uint) *big.Float {
	return amount.Quo(amount, big.NewFloat(math.Pow(10, float64(decimals))))
}
