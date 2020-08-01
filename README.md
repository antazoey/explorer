#  Vega Block Explorer
[![Netlify Status](https://api.netlify.com/api/v1/badges/4fff1f7f-315e-46dd-b821-2ddae569ebbc/deploy-status)](https://app.netlify.com/sites/vega-testnet-explorer/deploys)

This repo contains two things:
- A Go service that unpacks Vega transactions from Tendermint
- A client-side block explorer

## Deployment
Both the Go service & site are deployed to Netlify on build. Check out the Makefile to see what is involved there.

## Dependencies
The resulting site is *fragile* as it depends on:
- The Go service correctly decoding transactions;
- [Topgun Lookup](https://github.com/vegaprotocol/topgun-lookup), which in turn depends on [Topgun Service]()https://github.com/vegaprotocol/topgun-service
- [Vega](https://github.com/vegaprotocol/vega)'s GraphQL API
- [Vega](https://github.com/vegaprotocol/vega)'s REST API
- [Tendermint](https://github.com/tendermint/tendermint)'s RPC

So the chances of updates somewhere along the line breaking this are... pretty high!
