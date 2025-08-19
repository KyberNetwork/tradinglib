package finderengine

import (
	"context"
	"fmt"

	"github.com/KyberNetwork/tradinglib/pkg/finderengine/entity"
	"github.com/KyberNetwork/tradinglib/pkg/finderengine/finalizer"
	finderPkg "github.com/KyberNetwork/tradinglib/pkg/finderengine/finder"
	"github.com/pkg/errors"
)

type PathFinderEngine struct {
	finder    IFinder
	finalizer IFinalizer
}

func NewPathFinderEngine(
	finder IFinder,
	finalizer IFinalizer,
) *PathFinderEngine {
	return &PathFinderEngine{
		finder:    finder,
		finalizer: finalizer,
	}
}

func (p *PathFinderEngine) Find(ctx context.Context, params entity.FinderParams) (*entity.FinalizedRoute, error) {
	bestRoute, err := p.finder.Find(params)
	if err != nil {
		if errors.Is(err, finderPkg.ErrRouteNotFound) {
			return nil, fmt.Errorf("not found")
		}

		return nil, fmt.Errorf("failed to find route, err: %w", err)
	}

	if bestRoute.AMMBestRoute == nil {
		return nil, fmt.Errorf("invalid swap")
	}

	var ammRoute *entity.FinalizedRoute
	ammRoute, _ = finalizer.NewFinalizer().Finalize(params, bestRoute.AMMBestRoute)

	return ammRoute, fmt.Errorf("inval")
}

func (p *PathFinderEngine) SetFinder(finder IFinder) {
	p.finder = finder
}

func (p *PathFinderEngine) GetFinder() IFinder {
	return p.finder
}

func (p *PathFinderEngine) SetFinalizer(finalizer IFinalizer) {
	p.finalizer = finalizer
}

func (p *PathFinderEngine) GetFinalizer() IFinalizer {
	return p.finalizer
}
