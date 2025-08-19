package finderengine

import "github.com/KyberNetwork/tradinglib/pkg/finderengine/entity"

type IPathFinderEngine interface {
	Find(params entity.FinderParams) (*entity.FinalizedRoute, error)
	GetFinder() IFinder
	SetFinder(finder IFinder)
	GetFinalizer() IFinalizer
	SetFinalizer(finalizer IFinalizer)
}

type IFinder interface {
	Find(params entity.FinderParams) (*entity.BestRouteResult, error)
}

type IFinalizer interface {
	Finalize(params entity.FinderParams, bestRoute *entity.BestRouteResult) (*entity.FinalizedRoute, error)
}
