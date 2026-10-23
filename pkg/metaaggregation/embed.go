package metaaggregation

import _ "embed"

//go:embed MetaAggregationRouterV2.abi.json
var metaAggregationRouterV2JSON []byte

//go:embed KSAggregationRouterV3.abi.json
var ksAggregationRouterV3JSON []byte
