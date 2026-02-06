package flashblock

import (
	"context"
)

type DataSource string

const (
	NodeDataSource      DataSource = "node"
	BloxRouteDataSource DataSource = "bloxroute"
)

type Publisher interface {
	Publish(ctx context.Context, source DataSource, data Flashblock) error
}
