package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

type graphQueryBody struct {
	Query string `json:"query"`
}

type Position struct {
	Owner string `json:"owner"`
}

type GraphReplyV3 struct {
	Data struct {
		Positions []Position `json:"positions"`
	} `json:"data"`
}

type GraphReplyV3Snapshots struct {
	Data struct {
		Positions []Position `json:"positionsSnapshots"`
	} `json:"data"`
}

// GraphReplyV2pairs For finding id of a pair
type GraphReplyV2pairs struct {
	Data struct {
		Pair []struct {
			ID                     string `json:"id"`
			Reserve0               string `json:"reserve0"`
			Reserve1               string `json:"reserve1"`
			LiquidityProviderCount string `json:"liquidityProviderCount"`
			TotalSupply            string `json:"totalSupply"`
		} `json:"pairs"`
	} `json:"data"`
}

// GraphReplyV2positions get addresses of owners' positions
type GraphReplyV2positions struct {
	Data struct {
		LiquidityPositions []struct {
			LiquidityTokenBalance string `json:"liquidityTokenBalance"`
			User                  struct {
				ID string `json:"id"` //address
			} `json:"user"`
		} `json:"liquidityPositions"`
	} `json:"data"`
}

// GetOwnersV3 add owners to map
// supported pancake3 and uniswap3
func (app *App) GetOwnersV3(owners map[string]any, token0, token1 string) {
	var perPairLength int
	var perPairAddrs int
	for i := 0; ; i += 999 {
		reply := app.reqToV3(1000, uint64(i), token0, token1)
		perPairLength += len(reply.Data.Positions)
		for _, v := range reply.Data.Positions {
			if _, ok := owners[strings.ToLower(v.Owner)]; !ok {
				owners[strings.ToLower(v.Owner)] = nil
				perPairAddrs++
			}
		}
		if len(reply.Data.Positions) < 999 {
			fmt.Printf("Graph: %d unique addresses with %d positions for %s-%s\n", perPairAddrs, perPairLength, token0, token1)
			break
		}
	}
}

func (app *App) GetOwnersV3Snapshots(owners map[string]any, poolAddr string) {
	var perPairLength int
	var perPairAddrs int
	for i := 0; ; i += 999 {
		reply := app.reqToV3Snapshots(1000, uint64(i), poolAddr)
		perPairLength += len(reply.Data.Positions)
		for _, v := range reply.Data.Positions {
			if _, ok := owners[strings.ToLower(v.Owner)]; !ok {
				owners[strings.ToLower(v.Owner)] = nil
				perPairAddrs++
			}
		}
		if len(reply.Data.Positions) < 999 {
			fmt.Printf("Graph: %d unique addresses with %d positions for %s\n", perPairAddrs, perPairLength, poolAddr)
			break
		}
	}
}

// GetOwnersFromGraphV2 add owners to map
// supported pancake2 and uniswap2
func (app *App) GetOwnersFromGraphV2(owners map[string]any, token0Addr, token1Addr string) {
	var perPairLength int
	var perPairAddrs int
	pair := app.reqToV2pair(1, 0, token0Addr, token1Addr)
	for i := 0; ; i += 999 {
		reply := app.reqToV2Positions(1000, uint(i), pair.Data.Pair[0].ID)
		perPairLength += len(reply.Data.LiquidityPositions)
		for _, v := range reply.Data.LiquidityPositions {
			if _, ok := owners[strings.ToLower(v.User.ID)]; !ok {
				perPairAddrs++
				owners[strings.ToLower(v.User.ID)] = nil
			}
		}
		if len(reply.Data.LiquidityPositions) < 999 {
			fmt.Printf("thegraph.com: %d new addresses with %d positions for %s-%s\n", perPairAddrs, perPairLength, token0Addr, token1Addr)
			break
		}
	}
}

func (app *App) reqToV2pair(limit, skip uint, token0Addr, token1Addr string) GraphReplyV2pairs {
	queryStr := fmt.Sprintf(
		`{ pairs(first:%d, skip:%d, where:{and:[
	{token0_contains_nocase:"%s"},
	{token1_contains_nocase:"%s"}]}) {
	id
	reserve0
	reserve1
	totalSupply
	liquidityProviderCount
	token0 {
		id
		symbol
		decimals
		name
	}
	token1 {
		id
		symbol
		decimals
		name
	}}}`, limit, skip, token0Addr, token1Addr)
	resp, err := app.rawReq(queryStr)
	if err != nil {
		log.Println(err)
	}

	var reply GraphReplyV2pairs
	err = json.Unmarshal(resp, &reply)
	if err != nil {
		log.Fatalln("Can't unmarshal json: ", err)
	}
	//fmt.Println(reply)
	return reply
}

func (app *App) reqToV2Positions(limit, skip uint, pairID string) GraphReplyV2positions {
	queryStr := fmt.Sprintf(
		`{liquidityPositions(first:%d, skip:%d , where:{and:[{pair:"%s"}, 
      {liquidityTokenBalance_gt:0}]}) {
		liquidityTokenBalance
	  user {
		id
	}}}`, limit, skip, pairID)
	resp, err := app.rawReq(queryStr)
	var reply GraphReplyV2positions
	err = json.Unmarshal(resp, &reply)
	if err != nil {
		log.Fatalln("Can't unmarshal json: ", err)
	}
	return reply
}

func (app *App) reqToV3(limit, skip uint64, token0, token1 string) (reply GraphReplyV3) {
	queryStr := fmt.Sprintf("{positions(first:%d, skip:%d,where: {and: [{liquidity_gt:0}, {token0_:{symbol:\"%s\"}}, {token1_:{symbol:\"%s\"}}]	} ){owner}}", limit, skip, token0, token1)
	reply = GraphReplyV3{}
	resp, err := app.rawReq(queryStr)
	if err != nil {
		return GraphReplyV3{}
	}
	err = json.Unmarshal(resp, &reply)
	if err != nil {
		log.Fatalln("Can't unmarshal json: ", err)
	}
	return reply
}

func (app *App) reqToV3Snapshots(limit, skip uint64, poolAddr string) (reply GraphReplyV3Snapshots) {
	//queryStr := fmt.Sprintf("{positions(first:%d, skip:%d,where: {and: [{liquidity_gt:0}, {token0_:{symbol:\"%s\"}}, {token1_:{symbol:\"%s\"}}]	} ){owner}}", limit, skip, token0, token1)
	queryStr := fmt.Sprintf(`{
  positionSnapshots(first:%d,skip:%d, where:{pool_:{id:"%s"}})
  {
    owner
  }
}`, limit, skip, poolAddr)
	reply = GraphReplyV3Snapshots{}
	fmt.Println(queryStr)
	resp, err := app.rawReq(queryStr)
	if err != nil {
		return GraphReplyV3Snapshots{}
	}
	err = json.Unmarshal(resp, &reply)
	if err != nil {
		log.Fatalln("Can't unmarshal json: ", err)
	}
	return reply
}

func (app *App) rawReq(query string) (resp []byte, err error) {

	var body graphQueryBody
	body.Query = query
	marshal, err := json.Marshal(&body)
	if err != nil {
		return nil, err
	}

	ctx, cancelFunc := context.WithDeadline(context.Background(), time.Now().Add(time.Second*25))
	defer cancelFunc()
	req, err := http.NewRequestWithContext(ctx, "POST", app.GraphQLURL, bytes.NewBuffer(marshal))
	if err != nil {
		log.Printf("Can't create req to %s: %v \n", app.GraphQLURL, err)
		return nil, err
	}

	reply, err := http.DefaultClient.Do(req) // running actual request
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			log.Println("request timeout, try again or check graphql url in the config")
		}
		log.Fatalln("Can't do request:", err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Can't close body: ", err)
		}
	}(reply.Body)
	resp, err = io.ReadAll(reply.Body)
	if err != nil {
		log.Fatalln("Can't read body: ", err)
	}
	return resp, nil
}
