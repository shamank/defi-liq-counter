package main

import (
	"defi-liq-counter/pancake2/bsc/pancakeV2pair"
	"defi-liq-counter/pancake3/bsc/pancakeV3pool"
	uniswap2pair "defi-liq-counter/uniswap2/eth/uniswapV2pair"
	uniswapV3poolBsc "defi-liq-counter/uniswap3/bsc/uniswapV3pool"
	uniswapV3poolEth "defi-liq-counter/uniswap3/eth/uniswapV3pool"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"math/big"
	"strings"
	"sync"
	"time"
)

type pairV2 interface {
	GetReserves(opts *bind.CallOpts) (struct {
		Reserve0           *big.Int
		Reserve1           *big.Int
		BlockTimestampLast uint32
	}, error)
	TotalSupply(opts *bind.CallOpts) (*big.Int, error)
	BalanceOf(opts *bind.CallOpts, address common.Address) (*big.Int, error)
}

func (app *App) CountV2(owners map[string]any) map[string]*ResultV2 {

	var fullResult = map[string]*ResultV2{}
	for _, appPair := range app.Pairs {
		var err error
		var totalSupply *big.Int
		var pair pairV2
		var reserves struct {
			Reserve0           *big.Int
			Reserve1           *big.Int
			BlockTimestampLast uint32
		}
		switch app.DexApp {
		case Uniswap2eth:
			pair, err = uniswap2pair.NewUniswapV2pair(common.HexToAddress(appPair.PoolAddress), app.ethClient)
			if err != nil {
				log.Fatalln("can't init pool contract: ", err)
			}
		case Pancake2bsc:
			pair, err = pancakeV2pair.NewPancakeV2pair(common.HexToAddress(appPair.PoolAddress), app.ethClient)
			if err != nil {
				log.Fatalln("can't init pool contract: ", err)
			}
		default:
			panic("this dex_app not supported")
		}
		reserves, err = pair.GetReserves(app.opts)
		if err != nil {
			log.Fatalln("Can't GetReserves: ", err)
		}
		totalSupply, err = pair.TotalSupply(app.opts)
		if err != nil {
			log.Fatalln("Can't get TotalSupply: ", err)
		}
		var progressCount uint
		for addr := range owners {
			time.Sleep(300 * time.Millisecond)
			fmt.Printf("progress: %d of %d (%.2f%%) \n", progressCount, len(owners), float64(progressCount)/float64(len(owners))*100.0)
			progressCount++
			balance, err := pair.BalanceOf(app.opts, common.HexToAddress(addr))
			if err != nil {
				log.Println("Can't take balanceOf", err)
			}
			temp0 := big.NewInt(0).Mul(reserves.Reserve0, balance)
			temp1 := big.NewInt(0).Mul(reserves.Reserve1, balance)
			providerToken0 := temp0.Div(temp0, totalSupply)
			providerToken1 := temp1.Div(temp1, totalSupply)
			resV2, ok := fullResult[appPair.PoolAddress]
			if !ok || resV2 == nil {
				fullResult[appPair.PoolAddress] = initResultV2(appPair.Token0, appPair.Token1, appPair.PoolAddress)
			}
			//fullResult[appPair.PoolAddress].totalLiqFromGraph.Add(fullResult[appPair.PoolAddress].totalLiqFromGraph, liqBalanace)
			fullResult[appPair.PoolAddress].totalAmount0.Add(fullResult[appPair.PoolAddress].totalAmount0, big.NewFloat(0).SetInt(providerToken0))
			fullResult[appPair.PoolAddress].totalAmount1.Add(fullResult[appPair.PoolAddress].totalAmount1, big.NewFloat(0).SetInt(providerToken1))
			fullResult[appPair.PoolAddress].Amounts = append(fullResult[appPair.PoolAddress].Amounts, LiqPositionAmount{
				Address:             strings.ToLower(addr),
				Token0:              providerToken0.String(),
				Token1:              providerToken1.String(),
				Token0HumanReadable: humanReadableInt(providerToken0, appPair.Token0.Decimals).String(),
				Token1HumanReadable: humanReadableInt(providerToken1, appPair.Token1.Decimals).String(),
			})
		}
	}
	return fullResult
}

func initResultV2(token0, token1 Token, poolAddr string) *ResultV2 {
	return &ResultV2{
		totalAmount0:      new(big.Float),
		totalAmount1:      new(big.Float),
		totalLiqFromGraph: new(big.Float),
		Amounts:           make([]LiqPositionAmount, 0),
		Token0:            token0,
		Token1:            token1,
		PoolAddr:          poolAddr,
	}
}

func (app *App) InitV3(owners map[string]any) *V3 {
	v3 := new(V3)
	v3.Mutex = new(sync.Mutex)
	v3.result = map[string]*ResultV3{}
	for _, v := range app.Pairs {
		v3.result[v.Token0.Symbol+"/"+v.Token1.Symbol] = &ResultV3{
			totalAmount0:      new(big.Float),
			totalAmount1:      new(big.Float),
			tokensOwed0:       new(big.Int),
			totalFee:          new(big.Int),
			Token0:            v.Token0,
			Token1:            v.Token1,
			PoolAddr:          v.PoolAddress,
			Slot0SqrtPriceX96: new(big.Int),
			Amounts:           make([]LiqPositionAmount, 0),
		}
	}
	v3.owners = owners
	return v3
}

func (app *App) StartSyncV3(v3 *V3) {
	syncFunc := func() {
		for _, appPair := range app.Pairs {
			var slot0SqrtPriceX96 *big.Int
			switch app.DexApp {
			case Uniswap3eth:
				pool, err := uniswapV3poolEth.NewUniswapV3pool(common.HexToAddress(appPair.PoolAddress), app.ethClient)
				if err != nil {
					log.Println("slot0 can be 0 or outdated: can't init pool contract: ", err)
					continue
				}
				slot0, err := pool.Slot0(app.opts)
				if err != nil {
					log.Println("slot0 can be 0 or outdated: ", err)
					continue
				}
				slot0SqrtPriceX96 = slot0.SqrtPriceX96
			case Uniswap3bsc:
				pool, err := uniswapV3poolBsc.NewUniswapV3pool(common.HexToAddress(appPair.PoolAddress), app.ethClient)
				if err != nil {
					log.Println("slot0 can be 0 or outdated: can't init pool contract: ", err)
					continue
				}
				slot0, err := pool.Slot0(app.opts)
				if err != nil {
					log.Println("slot0 can be 0 or outdated: ", err)
					continue
				}
				slot0SqrtPriceX96 = slot0.SqrtPriceX96
			case Pancake3bsc:
				pool, err := pancakeV3pool.NewPancakeV3pool(common.HexToAddress(appPair.PoolAddress), app.ethClient)
				if err != nil {
					log.Println("slot0 can be 0 or outdated: can't init pool contract: ", err)
					continue
				}
				slot0, err := pool.Slot0(app.opts)
				if err != nil {
					log.Println("slot0 can be 0 or outdated: ", err)
					continue
				}
				slot0SqrtPriceX96 = slot0.SqrtPriceX96
			}
			v3.Lock()
			v3.result[appPair.Token0.Symbol+"/"+appPair.Token1.Symbol].Slot0SqrtPriceX96 = slot0SqrtPriceX96
			v3.Unlock()
		}
		fmt.Println("slot0 updated")
	}
	syncFunc() // sync first time to fill slot0
	ticker := time.NewTicker(5 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				syncFunc()
			}
		}
	}()
}

type V3 struct {
	result map[string]*ResultV3
	owners map[string]any
	*sync.Mutex
}

func (app *App) CountV3(v3 *V3) {

	var wait sync.WaitGroup
	var progressCount uint64
	var uniqPairs = map[string]struct {
		Token0      Token  `json:"token0"`
		Token1      Token  `json:"token1"`
		PoolAddress string `json:"pool_address"`
	}{}
	for _, appPair := range app.Pairs {
		uniqPairs[appPair.Token0.Address+"/"+appPair.Token1.Address] = struct {
			Token0      Token  `json:"token0"`
			Token1      Token  `json:"token1"`
			PoolAddress string `json:"pool_address"`
		}{Token0: appPair.Token0, Token1: appPair.Token1, PoolAddress: appPair.PoolAddress}
	}

	for addressStr := range v3.owners {
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("progress (owners): %d of %d (%.2f%%) \n", progressCount, len(v3.owners), float64(progressCount)/float64(len(v3.owners))*100.0)
		progressCount++
		var ownerAddress = common.HexToAddress(addressStr)

		wait.Add(1)
		go func(owner common.Address) {
			defer wait.Done()
			// time.Sleep(200 * time.Millisecond)
			//if strings.EqualFold(ownerAddress.Hex(), "0xE5b1BA172caA71A0D47997527524C4dbefeACF5a") {
			//	fmt.Println("0xE5b1BA172caA71A0D47997527524C4dbefeACF5a")
			//}
			positionsCount, err := app.posManager.BalanceOf(app.opts, ownerAddress)
			if err != nil {
				log.Printf("Can't get BalanceOf for address %s : %v \n", owner.Hex(), err)
				return
			}
			fmt.Println(positionsCount.String())

			var i int64
			fmt.Printf("%d positions for %s\n", positionsCount.Int64(), owner.Hex())
			for i = 0; i < positionsCount.Int64(); i++ {
				if i%20 == 0 && i > 20 {
					fmt.Printf("progress of %s: %d of %d (%.2f%%) \n", owner.Hex(), i, positionsCount.Int64(), float64(i)/float64(positionsCount.Int64())*100.0)
				}

				time.Sleep(100 * time.Millisecond)
				tokenID, err := app.posManager.TokenOfOwnerByIndex(app.opts, ownerAddress, big.NewInt(i))
				if err != nil {
					log.Printf("Can't get TokenOfOwnerByIndex for address %s : %v \n", owner.Hex(), err)
					continue
				}

				positionFromSmart, err := app.posManager.Positions(app.opts, tokenID)
				if err != nil {
					log.Printf("Can't get positions for address %s : %v \n", owner.Hex(), err)
					continue
				}

				for _, appPair := range uniqPairs {
					if strings.EqualFold(positionFromSmart.Token0.Hex(), appPair.Token0.Address) &&
						strings.EqualFold(positionFromSmart.Token1.Hex(), appPair.Token1.Address) {

						liqFloat := big.NewFloat(0).SetPrec(prec).SetInt(positionFromSmart.Liquidity)
						tickLowerFloat := big.NewFloat(0).SetPrec(prec).SetInt(positionFromSmart.TickLower)
						tickUpperFloat := big.NewFloat(0).SetPrec(prec).SetInt(positionFromSmart.TickUpper)

						v3.Lock()
						sqrtPriceX96Bigfloat := big.NewFloat(0).SetPrec(prec).SetInt(v3.result[appPair.Token0.Symbol+"/"+appPair.Token1.Symbol].Slot0SqrtPriceX96)
						v3.Unlock()

						amount0, amount1 := getTokenAmountsBigFloat(liqFloat, sqrtPriceX96Bigfloat, tickLowerFloat, tickUpperFloat)

						v3.Lock()
						v3.result[appPair.Token0.Symbol+"/"+appPair.Token1.Symbol].tokensOwed0 = v3.result[appPair.Token0.Symbol+"/"+appPair.Token1.Symbol].tokensOwed0.Add(v3.result[appPair.Token0.Symbol+"/"+appPair.Token1.Symbol].tokensOwed0, positionFromSmart.TokensOwed0)
						v3.result[appPair.Token0.Symbol+"/"+appPair.Token1.Symbol].totalFee = v3.result[appPair.Token0.Symbol+"/"+appPair.Token1.Symbol].totalFee.Add(v3.result[appPair.Token0.Symbol+"/"+appPair.Token1.Symbol].totalFee, positionFromSmart.Fee)

						amount0Safe := big.NewFloat(0).Set(amount0).SetPrec(prec)
						amount1Safe := big.NewFloat(0).Set(amount1).SetPrec(prec)
						if amount0.Cmp(big.NewFloat(0)) != 0 {
							v3.result[appPair.Token0.Symbol+"/"+appPair.Token1.Symbol].Amounts = append(v3.result[appPair.Token0.Symbol+"/"+appPair.Token1.Symbol].Amounts, LiqPositionAmount{
								Address:             strings.ToLower(ownerAddress.Hex()),
								Token0:              fmt.Sprintf("%.45f", amount0), // cogs or weis
								Token1:              fmt.Sprintf("%.45f", amount1), // cogs or weis
								Token0HumanReadable: humanReadableFloat(amount0, appPair.Token0.Decimals).String(),
								Token1HumanReadable: humanReadableFloat(amount1, appPair.Token1.Decimals).String(),
							})
						}

						v3.result[appPair.Token0.Symbol+"/"+appPair.Token1.Symbol].totalAmount0.Add(v3.result[appPair.Token0.Symbol+"/"+appPair.Token1.Symbol].totalAmount0, amount0Safe)
						v3.result[appPair.Token0.Symbol+"/"+appPair.Token1.Symbol].totalAmount1.Add(v3.result[appPair.Token0.Symbol+"/"+appPair.Token1.Symbol].totalAmount1, amount1Safe)
						v3.Unlock()
					}
				}
			}
		}(ownerAddress)
	}
	wait.Wait()
}
