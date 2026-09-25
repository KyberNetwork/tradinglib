package kipselipmm

const (
	ExchangeName = "kipseli-pmm"

	// DexType must stay in sync with kyberpmm.DexTypeKyberPMM in
	// github.com/KyberNetwork/kyberswap-dex-lib-private/pkg/liquidity-source/kyber-pmm.
	// Inlined as a literal so this module does not depend on that private repo.
	DexType = "kyber-pmm"
)
