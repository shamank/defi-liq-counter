## Counter of liquidity in evm dex apps

### Supported apps

|          | BSC | ETH | 
|----------|-----|-----|
| uniswap3 | ✔️  | ✔️  |   
| uniswap2 | -   | ✔️  |
| pancake3 | ✔️  | -   |
| pancake2 | ✔️  | -   |   

## Config

Config must be named as `config.json` (example will be below)

* dex_app (required) — name of dex with network, possible values:
(uniswap3-bsc, uniswap3-eth, pancake3-bsc, uniswap2-eth, pancake2-bsc);
* blockchain_url — endpoint of node/infura, examples:

    - https://bsc-mainnet.public.blastapi.io
    - https://mainnet.infura.io/v3/<API-KEY>

remember to turn on endpoints for BSC in infura if you're creating snapshots in BSC network.

* graphql_url — (required) url of provider for graphQL:
  ```https://gateway.thegraph.com/api/[api-key]/subgraphs/id/[subgraph-id]```


* pairs — array of structs:
    * pool_address — address of smart contract for pool (each pair of tokens has a unique address)
    * token0 (symbol, address and decimals)
    * token1 (symbol, address and decimals)

### Config example

```json
{
  "graphql_url": "https://gateway.thegraph.com/api/<API-KEY>/subgraphs/id/",
  "blockchain_url": "https://mainnet.infura.io/v3/<API-KEY>",
  "block_number": 0,
  "dex_app": "uniswap2-eth",
  "pairs": [
    {
      "token0": {
        "symbol": "",
        "address": "",
        "decimals": 18
      },
      "token1": {
        "symbol": "",
        "address": "",
        "decimals": 18
      },
      "pool_address": ""
    }
  ]
}
```


