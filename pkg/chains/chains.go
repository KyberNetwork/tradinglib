package chains

//go:generate enumer -type=ChainID -linecomment -json=true -sql=true -yaml
type ChainID int64 // nolint: recvcheck

const (
	Ethereum    ChainID = 1          // ethereum
	Optimism    ChainID = 10         // optimism
	Cronos      ChainID = 25         // cronos
	BSC         ChainID = 56         // bsc
	ETC         ChainID = 61         // etc
	Tomo        ChainID = 88         // tomo
	Gnosis      ChainID = 100        // gnosis
	Velas       ChainID = 106        // velas
	Polygon     ChainID = 137        // polygon
	BTTC        ChainID = 199        // bttc
	Fantom      ChainID = 250        // fantom
	KCC         ChainID = 321        // kcc
	Moonbeam    ChainID = 1284       // moonbeam
	Kava        ChainID = 2222       // kava
	Canto       ChainID = 7700       // canto
	Klaytn      ChainID = 8217       // klaytn
	EthereumPow ChainID = 10001      // ethw
	Fusion      ChainID = 32659      // fusion
	Arbitrum    ChainID = 42161      // arbitrum
	Celo        ChainID = 42220      // celo
	Oasis       ChainID = 42262      // oasis
	Avax        ChainID = 43114      // avax
	Aurora      ChainID = 1313161554 // aurora
)
