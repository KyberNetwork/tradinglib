# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What this repository is

`tradinglib` is a Go library, not a service.
It has no `main` package.
Downstream Kyber services consume it by version tag (current series `v0.11.x`).

## Two modules

The repository holds two separate Go modules.

| Path | Module | Notes |
| --- | --- | --- |
| `.` | `github.com/KyberNetwork/tradinglib` | go-ethereum v1.16.9 |
| `oneinch/` | `github.com/KyberNetwork/tradinglib/oneinch` | go-ethereum v1.16.2, `replace` directives for the Uniswap SDKs |

`oneinch/go.mod` requires `github.com/KyberNetwork/tradinglib v0.9.39`, a published tag.
It does not use a `replace` to the parent directory.
Two consequences matter for every change:

1. A root `go build ./...`, `go test ./...`, or `golangci-lint run` does not cover `oneinch/`.
   You must run each command a second time inside `oneinch/`.
2. An edit under `pkg/oneinch/` is invisible to `oneinch/pkg/encode` until a new tag ships.
   To test the edit locally, add a temporary `replace github.com/KyberNetwork/tradinglib => ../` and remove it before you commit.

`.github/workflows/ci.yml` never enters `oneinch/`, so that module gets no lint and no test signal in CI.

## Commands

Root module:

```sh
go build ./...
go test -race ./...                          # same flags as CI
go test ./pkg/mev/ -run TestUnmarshalSendBundleResponse1 -v   # single test
golangci-lint run --config=.golangci.yml     # golangci-lint v2.5.0
```

`oneinch/` module (run from the `oneinch` directory):

```sh
go build ./...
go test -race ./...
golangci-lint run --config=../.golangci.yml
```

The `oneinch/` lint command works, but it reports 5 pre-existing issues (3 `gosec`, 2 `modernize`) because CI never lints this module.

`.golangci.bck.yml` is a stale golangci-lint v1 config. No command uses it.

## Lint limits that block CI

`.golangci.yml` sets `default: all` with a disable list, so an unfamiliar linter can fail your change.
These limits are the ones that fail generated code most often:

- `funlen`: 80 lines, 50 statements.
- `cyclop`: max complexity 15.
- `lll`: line length.

`_test.go` files are exempt from `funlen`, `gocognit`, `lll`, and `cyclop`.
Formatters `gci`, `gofmt`, `gofumpt`, and `goimports` all run.
`pkg/basefee/op-geth/` is excluded from lint because it is vendored op-geth code.

## Tests that skip on purpose

Many tests start with a bare `t.Skip()`.
They are manual tests that need a live RPC endpoint, a builder endpoint, or a private key.
`pkg/mev/bundle_sender_test.go` alone holds 7 of them; `pkg/eth/trace_test.go` and the `pkg/flashblock` tests do the same.

Do not remove a `t.Skip()` to make a test run.
A green `go test ./...` gives no coverage of the bundle sender or listener paths.

## Architecture

### `pkg/mev` — bundle and transaction submission

One interface surface, many builder-specific clients.
`pkg.go` holds `IBundleSender`, `IBackrunSender`, `ISendRawTransaction`, `IGasBundleEstimator`, and the JSON-RPC method name constants.
Each builder gets its own client file: `bundle_sender.go` (generic `Client`), `blxr_bundle_sender.go`, `merkle_sender.go`, `mevshare_sender.go`, `layer2_sender.go`, `backrun_public_sender.go`, `bloxroute_backrunme_sender.go`.

Three parallel lists must stay in sync by hand when you add a builder:

1. The `BundleSenderType` enum in `pkg.go`.
2. The `Endpoint*` constants in `broadcaster.go`.
3. The `Builder*ID` constants in `broadcaster.go`.

`bundlesendertype_enumer.go` is generated. Regenerate it after you change the enum.

### `pkg/flashblock` — flashblock ingestion

Two data sources sit behind the single `Publisher` interface in `interface.go`:

- `NodeDataSource`: `rpc_listener.go` uses go-ethereum `EthSubscribe` on `newFlashblocks` and `pendingLogs`.
- `BloxRouteDataSource`: `blox_route_client.go` uses the bloXroute streamer API types in `block_route_types.go`.

### `pkg/oneinch` — 1inch order encoding and decoding

`decode.BytesIterator` is the shared primitive for sequential byte decoding.
The `limitorder` and `fusionorder` parsers are built on it, so a change to the iterator affects both.
`limitorder` is the largest area: `maker_traits.go`, `taker_traits.go`, `extension.go`, `fee_taker_extension.go`, and the settlement post-interaction encoders.

### `pkg/types` — JSON wire types

`BigInt`, `Bytes`, and `Duration` wrap standard types so that JSON marshals them as quoted strings.
They use mixed value and pointer receivers with `// nolint: recvcheck`.

### Other groups

- Chain and fee math: `pkg/chains`, `pkg/basefee` (with vendored op-geth EIP-1559 code for Base, BSC, and Polygon), `pkg/eth` (simulator, tracer, signature recovery, storage-slot overrides), `pkg/convert`.
- Calldata codecs: `pkg/metaaggregation` and `pkg/nativev2` decode router calldata from `go:embed`ed ABI JSON.
- Infrastructure helpers: `pkg/logging` (zap plus cclog plus Sentry), `pkg/metrics` (OpenTelemetry), `pkg/dbutil` (sqlx plus golang-migrate), `pkg/httpclient`, `pkg/httpsign`, `pkg/rate`, `pkg/bsync` (background local-state sync worker).
- Generic containers: `pkg/ds` (list, queue, stack), `pkg/hashset`, `x/syncmap`.
- `pkg/testutil` exports test helpers mainly for downstream services.
  Inside this repository only `oneinch/pkg/encode/lo1inch_test.go` imports it, for `NewBig10`, and it resolves that import from tradinglib v0.9.39 instead of the local source.
  No test in either module calls `MustNewDevelopmentDB`, so the postgres service in `ci.yml` is unused today.
  Do not set up a local database to run the tests.

## Code generation

- `enumer` generates `pkg/chains/chainid_enumer.go` and `pkg/mev/bundlesendertype_enumer.go`.
- The `go:generate` line in `pkg/mev/pkg.go` uses `-mod=vendor`, which cannot work because `vendor/` is gitignored.
  Use the path in its own comment instead: `go install github.com/dmarkham/enumer@latest`, then `enumer -type=BundleSenderType -linecomment`.
- ABI JSON is embedded, not generated: `pkg/metaaggregation/MetaAggregationRouterV2.abi.json` and `pkg/nativev2/nativev2.json`.
