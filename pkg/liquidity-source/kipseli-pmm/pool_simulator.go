package kipselipmm

import (
	"errors"
	"fmt"
	"maps"
	"math/big"
	"slices"

	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/pool"
	"github.com/KyberNetwork/tradinglib/pkg/convert"
)

type PoolSimulator struct {
	tokens                       map[string]Token
	tokenAddressesToTokenSymbols map[string]string
	pairs                        map[string]Pair // key: "base/quote"
	priorityQuotes               []string

	// key: base, this will use to keep track of sold/bought base amounts to avoid mega dumping.
	pairStates map[string]PairState
}

func NewPoolSimulator(pricingData PricingData, priorityQuotes []string) (*PoolSimulator, error) {
	if pricingData.IsZero() {
		return nil, errors.New("empty pricing data")
	}
	if len(priorityQuotes) == 0 {
		return nil, errors.New("empty priority quotes")
	}

	tokenAddressesToTokenSymbols := make(map[string]string, len(pricingData.Tokens))
	for _, token := range pricingData.Tokens {
		tokenAddressesToTokenSymbols[token.Address] = token.Symbol
	}

	pairs := make(map[string]Pair, len(pricingData.Pairs))
	for _, p := range pricingData.Pairs {
		key := pairKey(p.Base, p.Quote)
		pairs[key] = p
	}

	return &PoolSimulator{
		tokens:                       pricingData.Tokens,
		tokenAddressesToTokenSymbols: tokenAddressesToTokenSymbols,
		pairs:                        pairs,
		priorityQuotes:               priorityQuotes,
		pairStates:                   make(map[string]PairState),
	}, nil
}

func (p *PoolSimulator) CalcAmountOut(params pool.CalcAmountOutParams) (*pool.CalcAmountOutResult, error) {
	path, err := p.findPath(params.TokenAmountIn.Token, params.TokenOut)
	if err != nil {
		return nil, fmt.Errorf("find path: %w", err)
	}

	amountInF := convert.WeiToFloat(params.TokenAmountIn.Amount, path.TokenIn.Decimals)
	amountOutF, swapLogs, err := path.CalcAmountOut(amountInF)
	if err != nil {
		return nil, fmt.Errorf("calc amount out: %w", err)
	}

	amountOut, err := convert.FloatToWei(amountOutF, path.TokenOut.Decimals)
	if err != nil {
		return nil, fmt.Errorf("convert amount out: %w", err)
	}

	return &pool.CalcAmountOutResult{
		TokenAmountOut: &pool.TokenAmount{
			Token:  path.TokenOut.Address,
			Amount: amountOut,
		},
		Fee: &pool.TokenAmount{
			Token:  path.TokenIn.Address,
			Amount: big.NewInt(0),
		},
		SwapInfo: SwapExtra{
			TakerAsset:   path.TokenIn.Address,
			TakingAmount: params.TokenAmountIn.Amount.String(),
			MakerAsset:   path.TokenOut.Address,
			MakingAmount: amountOut.String(),
			IsExactOut:   false,
			SwapLogs:     swapLogs,
		},
	}, nil
}

func (p *PoolSimulator) CalcAmountIn(param pool.CalcAmountInParams) (*pool.CalcAmountInResult, error) {
	path, err := p.findPath(param.TokenIn, param.TokenAmountOut.Token)
	if err != nil {
		return nil, fmt.Errorf("find path: %w", err)
	}

	amountOutF := convert.WeiToFloat(param.TokenAmountOut.Amount, path.TokenOut.Decimals)
	amountInF, swapLogs, err := path.CalcAmountIn(amountOutF)
	if err != nil {
		return nil, fmt.Errorf("calc amount in: %w", err)
	}

	amountIn, err := convert.FloatToWei(amountInF, path.TokenIn.Decimals)
	if err != nil {
		return nil, fmt.Errorf("convert amount in: %w", err)
	}

	return &pool.CalcAmountInResult{
		TokenAmountIn: &pool.TokenAmount{
			Token:  path.TokenIn.Address,
			Amount: amountIn,
		},
		Fee: &pool.TokenAmount{
			Token:  path.TokenIn.Address,
			Amount: big.NewInt(0),
		},
		SwapInfo: SwapExtra{
			TakerAsset:   path.TokenIn.Address,
			TakingAmount: amountIn.String(),
			MakerAsset:   path.TokenOut.Address,
			MakingAmount: param.TokenAmountOut.Amount.String(),
			IsExactOut:   true,
			SwapLogs:     swapLogs,
		},
	}, nil
}

func (p *PoolSimulator) UpdateBalance(params pool.UpdateBalanceParams) {
	swapInfo, ok := params.SwapInfo.(SwapExtra)
	if !ok {
		panic(ErrInvalidSwapInfo)
	}

	for _, log := range swapInfo.SwapLogs {
		base := log.Base
		state := p.pairStates[base]
		state.SoldAmount += log.SoldBaseAmount
		state.BoughtAmount += log.BoughtBaseAmount
		p.pairStates[base] = state // update back
	}
}

func (p *PoolSimulator) findPath(tokenInAddr, tokenOutAddr string) (Path, error) {
	symbolIn, exists := p.tokenAddressesToTokenSymbols[tokenInAddr]
	if !exists {
		return Path{}, fmt.Errorf("%w: %s", ErrTokenNotFound, tokenInAddr)
	}
	symbolOut, exists := p.tokenAddressesToTokenSymbols[tokenOutAddr]
	if !exists {
		return Path{}, fmt.Errorf("%w: %s", ErrTokenNotFound, tokenOutAddr)
	}

	inToken := p.tokens[symbolIn]
	outToken := p.tokens[symbolOut]

	path := Path{
		TokenIn:  inToken,
		TokenOut: outToken,
	}

	// direct pair
	// try base/quote
	if pair, exists := p.getPair(symbolIn, symbolOut); exists {
		path = path.AddPair(false, pair, p.pairStates[pair.Base], inToken.MinTradeAmount)
		return path, nil
	}
	// try quote/base
	if pair, exists := p.getPair(symbolOut, symbolIn); exists {
		path = path.AddPair(true, pair, p.pairStates[pair.Base], outToken.MinTradeAmount)
		return path, nil
	}

	// single hop via priority quotes
	for _, midSymbol := range p.priorityQuotes {
		if midSymbol == symbolIn || midSymbol == symbolOut {
			continue
		}

		// try in/mid + out/mid
		pairIn, existsIn := p.getPair(symbolIn, midSymbol)
		pairOut, existsOut := p.getPair(symbolOut, midSymbol)
		if existsIn && existsOut {
			path = path.AddPair(false, pairIn, p.pairStates[pairIn.Base], inToken.MinTradeAmount)
			path = path.AddPair(true, pairOut, p.pairStates[pairOut.Base], outToken.MinTradeAmount)
			return path, nil
		}
	}

	return Path{}, fmt.Errorf("%w: %s -> %s", ErrNoPathFound, tokenInAddr, tokenOutAddr)
}

func (p *PoolSimulator) getPair(base, quote string) (Pair, bool) {
	key := pairKey(base, quote)
	pair, exists := p.pairs[key]
	return pair, exists
}

func (p *PoolSimulator) GetMetaInfo(_, _ string) any {
	return nil
}

func (p *PoolSimulator) CloneState() pool.IPoolSimulator {
	return &PoolSimulator{
		// shallow copy is enough, because we do not modify those maps/slices
		tokens:                       p.tokens,
		tokenAddressesToTokenSymbols: p.tokenAddressesToTokenSymbols,
		pairs:                        p.pairs,
		priorityQuotes:               p.priorityQuotes,
		// need to deep copy pairStates, because it will be modified
		pairStates: maps.Clone(p.pairStates),
	}
}

func (p *PoolSimulator) CanSwapTo(address string) []string {
	symbol, exists := p.tokenAddressesToTokenSymbols[address]
	if !exists {
		return nil
	}

	result := make([]string, 0, len(p.tokens))
	for _, tk := range p.tokens {
		if tk.Symbol == symbol {
			continue
		}

		result = append(result, tk.Address)
	}

	return result
}

func (p *PoolSimulator) CanSwapFrom(address string) []string {
	return p.CanSwapTo(address)
}

func (p *PoolSimulator) GetTokens() []string {
	tokens := make([]string, 0, len(p.tokens))
	for _, t := range p.tokens {
		tokens = append(tokens, t.Address)
	}

	slices.Sort(tokens)

	return tokens
}

func (p *PoolSimulator) GetReserves() []*big.Int {
	return nil
}

func (p *PoolSimulator) GetAddress() string {
	return "kipseli-kyber-pmm"
}

func (p *PoolSimulator) GetExchange() string {
	return ExchangeName
}

func (p *PoolSimulator) GetType() string {
	return DexType
}

func (p *PoolSimulator) GetTokenIndex(address string) int {
	tokens := p.GetTokens()
	return slices.Index(tokens, address)
}

func (p *PoolSimulator) CalculateLimit() map[string]*big.Int {
	return nil
}

func (p *PoolSimulator) GetApprovalAddress(_, _ string) string {
	return ""
}
