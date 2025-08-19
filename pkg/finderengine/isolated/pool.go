package isolated

import (
	"math/big"
	"sync"

	dexlibPool "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/pool"
)

type Pool struct {
	mu     sync.Mutex
	base   dexlibPool.IPoolSimulator
	local  dexlibPool.IPoolSimulator
	cloned bool
}

func NewIsolatedPool(base dexlibPool.IPoolSimulator) *Pool {
	return &Pool{
		base:   base,
		local:  base,
		cloned: false,
	}
}

func NewIsolatedPools(pools []dexlibPool.IPoolSimulator) []*Pool {
	isolatedPools := make([]*Pool, 0, len(pools))
	for i := range pools {
		isolatedPools = append(isolatedPools, NewIsolatedPool(pools[i]))
	}

	return isolatedPools
}

func (p *Pool) CalcAmountOut(params dexlibPool.CalcAmountOutParams) (*dexlibPool.CalcAmountOutResult, error) {
	return p.local.CalcAmountOut(params)
}

func (p *Pool) UpdateBalance(params dexlibPool.UpdateBalanceParams) {
	p.ensureClone()
	p.local.UpdateBalance(params)
}

func (p *Pool) CloneState() dexlibPool.IPoolSimulator {
	src := p.local
	return &Pool{base: src, local: src.CloneState(), cloned: true}
}

func (p *Pool) Reset() {
	p.mu.Lock()
	p.local = p.base
	p.cloned = false
	p.mu.Unlock()
}

func (p *Pool) CanSwapFrom(address string) []string {
	return p.local.CanSwapFrom(address)
}
func (p *Pool) GetTokens() []string                 { return p.local.GetTokens() }
func (p *Pool) GetReserves() []*big.Int             { return p.local.GetReserves() }
func (p *Pool) GetAddress() string                  { return p.local.GetAddress() }
func (p *Pool) GetExchange() string                 { return p.local.GetExchange() }
func (p *Pool) GetType() string                     { return p.local.GetType() }
func (p *Pool) GetTokenIndex(address string) int    { return p.local.GetTokenIndex(address) }
func (p *Pool) CalculateLimit() map[string]*big.Int { return p.local.CalculateLimit() }
func (p *Pool) CanSwapTo(address string) []string   { return p.local.CanSwapTo(address) }
func (p *Pool) GetMetaInfo(tokenIn, tokenOut string) any {
	return p.local.GetMetaInfo(tokenIn, tokenOut)
}

func (p *Pool) ensureClone() {
	if p.cloned {
		return
	}
	p.mu.Lock()
	if !p.cloned {
		p.clone()
	}
	p.mu.Unlock()
}

func (p *Pool) clone() {
	p.local = p.base.CloneState()
	p.cloned = true
}
