package poolsimulators

import (
	"errors"
	"fmt"
	"math/big"

	ksent "github.com/KyberNetwork/kyberswap-dex-lib/pkg/entity"
	balancerv1 "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/balancer-v1"
	balancerv2composablestable "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/balancer-v2/composable-stable"
	balancerv2stable "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/balancer-v2/stable"
	balancerv2weighted "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/balancer-v2/weighted"
	bancorv21 "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/bancor-v21"
	bancorv3 "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/bancor-v3"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/bebop"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/bedrock/unieth"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/clipper"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/curve/plain"
	stableng "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/curve/stable-ng"
	tricryptong "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/curve/tricrypto-ng"
	twocryptong "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/curve/twocrypto-ng"
	daiusds "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/dai-usds"
	deltaswapv1 "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/deltaswap-v1"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/dexalot"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/dodo/classical"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/dodo/dpp"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/dodo/dsp"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/dodo/dvm"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/ethena/susde"
	ethervista "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/ether-vista"
	etherfieeth "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/etherfi/eeth"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/etherfi/weeth"
	dexT1 "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/fluid/dex-t1"
	vaultT1 "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/fluid/vault-t1"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/frax/sfrxeth"
	sfrxeth_convertor "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/frax/sfrxeth-convertor"
	generic_simple_rate "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/generic-simple-rate"
	gyro2clp "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/gyroscope/2clp"
	gyro3clp "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/gyroscope/3clp"
	gyroeclp "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/gyroscope/eclp"
	hashflowv3 "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/hashflow-v3"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/integral"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/kelp/rseth"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/litepsm"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/maker/savingsdai"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/mantle/meth"
	mkrsky "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/mkr-sky"
	nativev1 "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/native-v1"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/nomiswap/nomiswapstable"
	ondo_usdy "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/ondo-usdy"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/primeeth"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/puffer/pufeth"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/renzo/ezeth"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/ringswap"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/rocketpool/reth"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/staderethx"
	swaapv2 "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/swaap-v2"
	swellrsweth "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/swell/rsweth"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/swell/sweth"
	uniswapv1 "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/uniswap-v1"
	uniswapv2 "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/uniswap-v2"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/usd0pp"
	velocorev2cpmm "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/velocore-v2/cpmm"
	velocorev2wombatstable "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/velocore-v2/wombat-stable"
	velodromev1 "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/velodrome-v1"
	velodromev2 "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/velodrome-v2"
	woofiv21 "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/woofi-v21"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/pooltypes"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/algebrav1"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/balancer"
	balancercomposablestable "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/balancer-composable-stable"
	balancerstable "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/balancer/stable"
	balancerweighted "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/balancer/weighted"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/camelot"
	curveaave "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/curve/aave"
	curvebase "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/curve/base"
	curvecompound "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/curve/compound"
	curveplainoracle "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/curve/plain-oracle"
	curvetricrypto "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/curve/tricrypto"
	curvetwo "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/curve/two"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/dmm"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/elastic"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/equalizer"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/fraxswap"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/fulcrom"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/fxdx"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/gmx"
	gmxglp "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/gmx-glp"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/iziswap"
	kokonutcrypto "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/kokonut-crypto"
	kyberpmm "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/kyber-pmm"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/lido"
	lidosteth "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/lido-steth"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/limitorder"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/liquiditybookv20"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/liquiditybookv21"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/madmex"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/makerpsm"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/mantisswap"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/maverickv1"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/metavault"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/nuriv2"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/pancakev3"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/platypus"
	polmatic "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/pol-matic"
	pkgpool "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/pool"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/quickperps"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/ramsesv2"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/saddle"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/slipstream"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/smardex"
	solidlyv3 "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/solidly-v3"
	swapbasedperp "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/swapbased-perp"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/syncswap/syncswapclassic"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/syncswap/syncswapstable"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/synthetix"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/uniswap"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/uniswapv3"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/usdfi"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/velocimeter"
	velodromelegacy "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/velodrome"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/vooi"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/wombat/wombatlsd"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/wombat/wombatmain"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/woofiv2"
	zkera "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/zkera-finance"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/swaplimit"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/valueobject"
)

var ErrPoolTypeNotSupported = errors.New("pool type is not supported")

// PoolSimulatorFromPool
// nolint: funlen, gocyclo, cyclop, maintidx
func PoolSimulatorFromPool(pool ksent.Pool, chainID uint) (pkgpool.IPoolSimulator, error) {
	var (
		pSim pkgpool.IPoolSimulator
		err  error
	)

	switch pool.Type {
	case pooltypes.PoolTypes.Uni, pooltypes.PoolTypes.Firebird,
		pooltypes.PoolTypes.Biswap, pooltypes.PoolTypes.Polydex:
		pSim, err = uniswap.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.UniswapV2:
		pSim, err = uniswapv2.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.UniswapV3:
		pSim, err = uniswapv3.NewPoolSimulator(pool, valueobject.ChainID(chainID))
	case pooltypes.PoolTypes.Dmm:
		pSim, err = dmm.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.Elastic:
		pSim, err = elastic.NewPoolSimulator(pool, valueobject.ChainID(chainID))
	case pooltypes.PoolTypes.BalancerV1:
		pSim, err = balancerv1.NewPoolSimulator(pool)
	case string(balancer.DexTypeBalancerStable):
		pSim, err = balancerstable.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.BalancerV2Stable:
		pSim, err = balancerv2stable.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.LegacyBalancerWeighted:
		pSim, err = balancerweighted.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.BalancerV2Weighted:
		pSim, err = balancerv2weighted.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.LegacyBalancerComposableStable:
		pSim, err = balancercomposablestable.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.BalancerV2ComposableStable:
		pSim, err = balancerv2composablestable.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.DodoClassical:
		pSim, err = classical.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.DodoVendingMachine:
		pSim, err = dvm.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.DodoStablePool:
		pSim, err = dsp.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.DodoPrivatePool:
		pSim, err = dpp.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.CurveStablePlain:
		pSim, err = plain.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.CurveStableNg:
		pSim, err = stableng.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.CurveBase:
		pSim, err = curvebase.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.CurveTwo:
		pSim, err = curvetwo.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.CurveTwoCryptoNg:
		pSim, err = twocryptong.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.CurveAave:
		pSim, err = curveaave.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.CurveCompound:
		pSim, err = curvecompound.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.CurveTricrypto:
		pSim, err = curvetricrypto.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.CurveTriCryptoNg:
		pSim, err = tricryptong.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.CurvePlainOracle:
		pSim, err = curveplainoracle.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.Saddle, pooltypes.PoolTypes.Nerve,
		pooltypes.PoolTypes.OneSwap, pooltypes.PoolTypes.IronStable:
		pSim, err = saddle.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.Velodrome:
		pSim, err = velodromev1.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.Ramses, pooltypes.PoolTypes.MuteSwitch,
		pooltypes.PoolTypes.Dystopia, pooltypes.PoolTypes.Pearl:
		pSim, err = velodromelegacy.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.VelodromeV2:
		pSim, err = velodromev2.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.Velocimeter:
		pSim, err = velocimeter.NewPool(pool)
	case pooltypes.PoolTypes.PlatypusBase, pooltypes.PoolTypes.PlatypusPure, pooltypes.PoolTypes.PlatypusAvax:
		pSim, err = platypus.NewPoolSimulator(pool, valueobject.ChainID(chainID))
	case pooltypes.PoolTypes.GMX:
		pSim, err = gmx.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.MakerPSM:
		pSim, err = makerpsm.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.Synthetix:
		pSim, err = synthetix.NewPoolSimulator(pool, valueobject.ChainID(chainID))
	case pooltypes.PoolTypes.MadMex:
		pSim, err = madmex.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.Metavault:
		pSim, err = metavault.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.Lido:
		pSim, err = lido.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.LidoStEth:
		pSim, err = lidosteth.NewPoolSimulator(pool, valueobject.ChainID(chainID))
	case pooltypes.PoolTypes.Fraxswap:
		pSim, err = fraxswap.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.Camelot:
		pSim, err = camelot.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.LimitOrder:
		pSim, err = limitorder.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.SyncSwapClassic:
		pSim, err = syncswapclassic.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.SyncSwapStable:
		pSim, err = syncswapstable.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.PancakeV3:
		pSim, err = pancakev3.NewPoolSimulator(pool, valueobject.ChainID(chainID))
	case pooltypes.PoolTypes.MaverickV1:
		pSim, err = maverickv1.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.SolidlyV3:
		pSim, err = solidlyv3.NewPoolSimulator(pool, valueobject.ChainID(chainID))
	case pooltypes.PoolTypes.WombatMain:
		pSim, err = wombatmain.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.WombatLsd:
		pSim, err = wombatlsd.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.Smardex:
		pSim, err = smardex.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.SwapBasedPerp:
		pSim, err = swapbasedperp.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.AlgebraV1:
		pSim, err = algebrav1.NewPoolSimulator(pool, DefaultGasAlgebra[valueobject.Exchange(pool.Exchange)])
	case pooltypes.PoolTypes.IZiSwap:
		pSim, err = iziswap.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.GMXGLP:
		pSim, err = gmxglp.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.RamsesV2:
		pSim, err = ramsesv2.NewPoolSimulator(pool, valueobject.ChainID(chainID))
	case pooltypes.PoolTypes.USDFi:
		pSim, err = usdfi.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.Equalizer:
		pSim, err = equalizer.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.MantisSwap:
		pSim, err = mantisswap.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.KokonutCrypto:
		pSim, err = kokonutcrypto.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.PolMatic:
		pSim, err = polmatic.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.Fxdx:
		pSim, err = fxdx.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.Vooi:
		pSim, err = vooi.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.QuickPerps:
		pSim, err = quickperps.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.VelocoreV2CPMM:
		pSim, err = velocorev2cpmm.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.VelocoreV2WombatStable:
		pSim, err = velocorev2wombatstable.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.Fulcrom:
		pSim, err = fulcrom.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.Gyroscope2CLP:
		pSim, err = gyro2clp.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.Gyroscope3CLP:
		pSim, err = gyro3clp.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.GyroscopeECLP:
		pSim, err = gyroeclp.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.ZkEraFinance:
		pSim, err = zkera.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.BancorV3:
		pSim, err = bancorv3.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.EtherfiWEETH:
		pSim, err = weeth.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.EtherfiEETH:
		pSim, err = etherfieeth.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.KelpRSETH:
		pSim, err = rseth.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.RocketPoolRETH:
		pSim, err = reth.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.EthenaSusde:
		pSim, err = susde.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.MakerSavingsDai:
		pSim, err = savingsdai.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.HashflowV3:
		pSim, err = hashflowv3.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.KyberPMM:
		pSim, err = kyberpmm.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.NativeV1:
		pSim, err = nativev1.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.SwaapV2:
		pSim, err = swaapv2.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.EtherVista:
		pSim, err = ethervista.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.LitePSM:
		pSim, err = litepsm.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.Integral:
		pSim, err = integral.NewPoolSimulator(pool)
	case bancorv21.DexType:
		pSim, err = bancorv21.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.RenzoEZETH:
		pSim, err = ezeth.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.BedrockUniETH:
		pSim, err = unieth.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.PufferPufETH:
		pSim, err = pufeth.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.SwellSWETH:
		pSim, err = sweth.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.SwellRSWETH:
		pSim, err = swellrsweth.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.Slipstream:
		pSim, err = slipstream.NewPoolSimulator(pool, valueobject.ChainID(chainID))
	case pooltypes.PoolTypes.NuriV2:
		pSim, err = nuriv2.NewPoolSimulator(pool, valueobject.ChainID(chainID))
	case pooltypes.PoolTypes.Usd0PP:
		pSim, err = usd0pp.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.FluidVaultT1:
		pSim, err = vaultT1.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.MkrSky:
		pSim, err = mkrsky.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.DaiUsds:
		pSim, err = daiusds.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.LiquidityBookV20:
		pSim, err = liquiditybookv20.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.LiquidityBookV21:
		pSim, err = liquiditybookv21.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.Bebop:
		pSim, err = bebop.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.UniswapV1:
		pSim, err = uniswapv1.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.Dexalot:
		pSim, err = dexalot.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.FluidDexT1:
		pSim, err = dexT1.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.GenericSimpleRate:
		pSim, err = generic_simple_rate.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.PrimeETH:
		pSim, err = primeeth.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.StaderETHx:
		pSim, err = staderethx.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.MantleETH:
		pSim, err = meth.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.WooFiV2:
		pSim, err = woofiv2.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.WooFiV21:
		pSim, err = woofiv21.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.DeltaSwapV1:
		pSim, err = deltaswapv1.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.NomiSwapStable:
		pSim, err = nomiswapstable.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.RingSwap:
		pSim, err = ringswap.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.OndoUSDY:
		pSim, err = ondo_usdy.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.Clipper:
		pSim, err = clipper.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.SfrxETH:
		pSim, err = sfrxeth.NewPoolSimulator(pool)
	case pooltypes.PoolTypes.SfrxETHConvertor:
		pSim, err = sfrxeth_convertor.NewPoolSimulator(pool)
	default:
		err = fmt.Errorf("%w: %s %s", ErrPoolTypeNotSupported, pool.Type, pool.Address)
	}

	return pSim, err
}

func newSwapLimit(dex string, limit map[string]*big.Int) pkgpool.SwapLimit {
	switch dex {
	case pooltypes.PoolTypes.Synthetix,
		pooltypes.PoolTypes.LimitOrder,
		pooltypes.PoolTypes.KyberPMM,
		pooltypes.PoolTypes.NativeV1,
		pooltypes.PoolTypes.Dexalot:
		return swaplimit.NewInventory(dex, limit)
	}

	return nil
}

func NewSwapLimit(limits map[string]map[string]*big.Int) map[string]pkgpool.SwapLimit {
	limitMap := make(map[string]pkgpool.SwapLimit, len(limits))

	for dex, limit := range limits {
		limitMap[dex] = newSwapLimit(dex, limit)
	}

	return limitMap
}
