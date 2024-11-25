module github.com/KyberNetwork/tradinglib

go 1.22.7

replace (
	github.com/daoleno/uniswap-sdk-core v0.1.5 => github.com/KyberNetwork/uniswap-sdk-core v0.1.5
	github.com/daoleno/uniswapv3-sdk v0.4.0 => github.com/KyberNetwork/uniswapv3-sdk v0.5.0
)

require (
	github.com/KyberNetwork/kyber-trace-go v0.1.1
	github.com/KyberNetwork/kyberswap-dex-lib v0.80.8
	github.com/duoxehyon/mev-share-go v0.3.0
	github.com/ethereum/go-ethereum v1.13.14
	github.com/flashbots/mev-share-node v0.0.0-20230926173018-7862d944990a
	github.com/golang-migrate/migrate/v4 v4.17.0
	github.com/google/uuid v1.6.0
	github.com/holiman/uint256 v1.3.1
	github.com/jmoiron/sqlx v1.3.5
	github.com/shopspring/decimal v1.4.0
	github.com/sourcegraph/conc v0.3.0
	github.com/stretchr/testify v1.9.0
	github.com/test-go/testify v1.1.4
	go.opentelemetry.io/otel/metric v1.22.0
	golang.org/x/exp v0.0.0-20240909161429-701f63a606c0
)

require (
	github.com/KyberNetwork/blockchain-toolkit v0.8.1 // indirect
	github.com/KyberNetwork/elastic-go-sdk/v2 v2.0.2 // indirect
	github.com/KyberNetwork/ethrpc v0.7.3 // indirect
	github.com/KyberNetwork/iZiSwap-SDK-go v1.1.0 // indirect
	github.com/KyberNetwork/int256 v0.1.4 // indirect
	github.com/KyberNetwork/logger v0.2.1 // indirect
	github.com/KyberNetwork/pancake-v3-sdk v0.2.0 // indirect
	github.com/KyberNetwork/uniswapv3-sdk-uint256 v0.5.0 // indirect
	github.com/Microsoft/go-winio v0.6.2 // indirect
	github.com/bits-and-blooms/bitset v1.14.3 // indirect
	github.com/btcsuite/btcd/btcec/v2 v2.3.4 // indirect
	github.com/btcsuite/btcd/chaincfg/chainhash v1.0.2 // indirect
	github.com/cenkalti/backoff/v4 v4.2.1 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/consensys/bavard v0.1.15 // indirect
	github.com/consensys/gnark-crypto v0.12.1 // indirect
	github.com/crate-crypto/go-kzg-4844 v0.7.0 // indirect
	github.com/daoleno/uniswap-sdk-core v0.1.7 // indirect
	github.com/daoleno/uniswapv3-sdk v0.4.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/deckarep/golang-set/v2 v2.6.0 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.3.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/ethereum/c-kzg-4844 v1.0.3 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-ole/go-ole v1.3.0 // indirect
	github.com/go-resty/resty/v2 v2.14.0 // indirect
	github.com/goccy/go-json v0.10.3 // indirect
	github.com/golang/mock v1.6.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/golang/snappy v0.0.5-0.20220116011046-fa5810519dcb // indirect
	github.com/gorilla/websocket v1.5.3 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.18.1 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/huin/goupnp v1.3.0 // indirect
	github.com/jackpal/go-nat-pmp v1.0.2 // indirect
	github.com/lib/pq v1.10.9 // indirect
	github.com/machinebox/graphql v0.2.2 // indirect
	github.com/metachris/flashbotsrpc v0.6.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mmcloughlin/addchain v0.4.0 // indirect
	github.com/orcaman/concurrent-map v1.0.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/redis/go-redis/v9 v9.0.2 // indirect
	github.com/samber/lo v1.38.1 // indirect
	github.com/shirou/gopsutil v3.21.11+incompatible // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/supranational/blst v0.3.13 // indirect
	github.com/syndtr/goleveldb v1.0.1-0.20210819022825-2ae1ddf74ef7 // indirect
	github.com/tklauser/go-sysconf v0.3.14 // indirect
	github.com/tklauser/numcpus v0.8.0 // indirect
	github.com/ybbus/jsonrpc/v3 v3.1.4 // indirect
	github.com/yusufpapurcu/wmi v1.2.4 // indirect
	go.opentelemetry.io/otel v1.22.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric v0.42.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc v0.42.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp v0.42.0 // indirect
	go.opentelemetry.io/otel/sdk v1.22.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.19.0 // indirect
	go.opentelemetry.io/otel/trace v1.22.0 // indirect
	go.opentelemetry.io/proto/otlp v1.0.0 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.27.0 // indirect
	golang.org/x/crypto v0.27.0 // indirect
	golang.org/x/net v0.29.0 // indirect
	golang.org/x/sync v0.8.0 // indirect
	golang.org/x/sys v0.25.0 // indirect
	golang.org/x/text v0.18.0 // indirect
	golang.org/x/time v0.6.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240123012728-ef4313101c80 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240123012728-ef4313101c80 // indirect
	google.golang.org/grpc v1.62.0 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.2.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	rsc.io/tmplfunc v0.0.3 // indirect
)
