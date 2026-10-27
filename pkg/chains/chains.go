package chains

//go:generate enumer -type=ChainID -linecomment -json=true -sql=true -yaml
type ChainID int64 // nolint: recvcheck

const (
	Ethereum     ChainID = 1          // ethereum
	Optimism     ChainID = 10         // optimism
	Cronos       ChainID = 25         // cronos
	BSC          ChainID = 56         // bsc
	ETC          ChainID = 61         // etc
	Tomo         ChainID = 88         // tomo
	Gnosis       ChainID = 100        // gnosis
	Velas        ChainID = 106        // velas
	Unichain     ChainID = 130        // unichain
	Polygon      ChainID = 137        // polygon
	Sonic        ChainID = 146        // sonic
	Manta        ChainID = 169        // manta
	BTTC         ChainID = 199        // bttc
	OpBNB        ChainID = 204        // opbnb
	Fantom       ChainID = 250        // fantom
	Fraxtal      ChainID = 252        // fraxtal
	KCC          ChainID = 321        // kcc
	ZkSync       ChainID = 324        // zksync
	Near         ChainID = 397        // near
	IntentNear   ChainID = 398        // intentnear
	WorldChain   ChainID = 480        // worldchain
	HyperEVM     ChainID = 999        // hyperevm
	Metis        ChainID = 1088       // metis
	PolygonZkEVM ChainID = 1101       // polygonzkevm
	Moonbeam     ChainID = 1284       // moonbeam
	Moonriver    ChainID = 1285       // moonriver
	Sei          ChainID = 1329       // sei
	Soneium      ChainID = 1868       // soneium
	Ronin        ChainID = 2020       // ronin
	Kava         ChainID = 2222       // kava
	Abstract     ChainID = 2741       // abstract
	Robinhood    ChainID = 4663       // robinhood
	Mantle       ChainID = 5000       // mantle
	Canto        ChainID = 7700       // canto
	Klaytn       ChainID = 8217       // klaytn
	Base         ChainID = 8453       // base
	EthereumPow  ChainID = 10001      // ethw
	Fusion       ChainID = 32659      // fusion
	Mode         ChainID = 34443      // mode
	Arbitrum     ChainID = 42161      // arbitrum
	Celo         ChainID = 42220      // celo
	Oasis        ChainID = 42262      // oasis
	Avax         ChainID = 43114      // avax
	Ink          ChainID = 57073      // ink
	Linea        ChainID = 59144      // linea
	Berachain    ChainID = 80094      // berachain
	Blast        ChainID = 81457      // blast
	Taiko        ChainID = 167000     // taiko
	Scroll       ChainID = 534352     // scroll
	Zora         ChainID = 7777777    // zora
	Aurora       ChainID = 1313161554 // aurora
)
