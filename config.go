package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/url"
	"os"
)

type DexApp string

const (
	Uniswap3bsc DexApp = "uniswap3-bsc"
	Uniswap3eth DexApp = "uniswap3-eth"
	Uniswap2eth DexApp = "uniswap2-eth"
	Pancake3bsc DexApp = "pancake3-bsc"
	Pancake2bsc DexApp = "pancake2-bsc"
)

func (dexApp DexApp) String() string {
	return string(dexApp)
}

func (dexApp DexApp) Validate() bool {
	validDexApps := map[DexApp]struct{}{
		Uniswap3bsc: {},
		Uniswap3eth: {},
		Uniswap2eth: {},
		Pancake3bsc: {},
		Pancake2bsc: {},
	}
	_, isValid := validDexApps[dexApp]
	return isValid
}

type Token struct {
	Symbol   string `json:"symbol"`
	Address  string `json:"address"`
	Decimals uint   `json:"decimals"`
}

type Pair struct {
	Token0      Token  `json:"token0"`
	Token1      Token  `json:"token1"`
	PoolAddress string `json:"pool_address"`
}

type Config struct {
	Pairs         []Pair `json:"pairs"`
	DexApp        DexApp `json:"dex_app"`
	GraphQLURL    string `json:"graphQL_url"`
	BlockchainURL string `json:"blockchain_url"`
	BlockNumber   uint   `json:"block_number"`
}

func loadConfig(filePath string) (Config, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, errors.New("unable to read configuration file: " + err.Error())
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return Config{}, errors.New("failed to parse configuration: " + err.Error())
	}

	return config, nil
}

func validateConfig(config Config) error {
	if !config.DexApp.Validate() {
		return errors.New("invalid dex_app, refer to README.md")
	}

	if config.BlockchainURL == "" {
		return errors.New("blockchain URL cannot be empty")
	}

	if _, err := url.ParseRequestURI(config.GraphQLURL); err != nil {
		return errors.New("invalid GraphQL URL: " + err.Error())
	}

	return nil
}

func newAppFromConf() *App {
	config, err := loadConfig("config.json")
	if err != nil {
		log.Fatalln(err)
	}

	if err := validateConfig(config); err != nil {
		log.Fatalln(err)
	}

	return &App{Config: config}
}
