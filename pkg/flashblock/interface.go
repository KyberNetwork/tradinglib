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
	PublishFlashBlock(ctx context.Context, source DataSource, data Flashblock) error
}
