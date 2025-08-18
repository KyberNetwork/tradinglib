package isolatedpools

import (
	"math/big"

	dexlibPool "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/pool"
)

type IsolatedPool struct {
	base   dexlibPool.IPoolSimulator
	local  dexlibPool.IPoolSimulator
	cloned bool
}

func NewIsolatedPools(base dexlibPool.IPoolSimulator) *IsolatedPool {
	return &IsolatedPool{
		base:   base,
		local:  base,
		cloned: false,
	}
}

func (p *IsolatedPool) CalcAmountOut(params dexlibPool.CalcAmountOutParams) (*dexlibPool.CalcAmountOutResult, error) {
	return p.local.CalcAmountOut(params)
}

func (p *IsolatedPool) UpdateBalance(params dexlibPool.UpdateBalanceParams) {
	p.ensureClone()
	p.local.UpdateBalance(params)
}

func (p *IsolatedPool) CloneState() *IsolatedPool {
	return nil
}

func (p *IsolatedPool) CanSwapFrom(address string) []string {
	return p.local.CanSwapFrom(address)
}
func (p *IsolatedPool) GetTokens() []string     { return p.local.GetTokens() }
func (p *IsolatedPool) GetReserves() []*big.Int { return p.local.GetReserves() }
func (p *IsolatedPool) GetAddress() string      { return p.local.GetAddress() }
func (p *IsolatedPool) GetExchange() string     { return p.local.GetExchange() }
func (p *IsolatedPool) GetType() string         { return p.local.GetType() }
func (p *IsolatedPool) GetMetaInfo(tokenIn, tokenOut string) any {
	return p.local.GetMetaInfo(tokenIn, tokenOut)
}
func (p *IsolatedPool) GetTokenIndex(address string) int    { return p.local.GetTokenIndex(address) }
func (p *IsolatedPool) CalculateLimit() map[string]*big.Int { return p.local.CalculateLimit() }
func (p *IsolatedPool) CanSwapTo(address string) []string   { return p.local.CanSwapTo(address) }

func (p *IsolatedPool) ensureClone() {
	if p.cloned {
		return
	}

	p.clone()
}

func (p *IsolatedPool) clone() {
	p.local = p.base.CloneState()
	p.cloned = true
}
