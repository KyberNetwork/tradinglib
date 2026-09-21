package dexrouterwrapper

import _ "embed"

//go:embed DexRouterWrapper.abi.json
var wrapperABIJSON []byte

//go:embed DexRouterWrapper.bin
var wrapperCreationBytecodeHex string

//go:embed DexRouterWrapper.bin-runtime
var wrapperDeployedBytecodeHex string
