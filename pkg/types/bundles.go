package types

import (
	"time"

	"github.com/lib/pq"
)

type ExecuteBundle struct {
	ID          uint64         `db:"id"`
	TxHashes    pq.StringArray `db:"tx_hashes"` // Use pq.StringArray instead of []string
	BundleHash  string         `db:"bundle_hash"`
	BlockNumber uint64         `db:"block_number"`
	SubmitAt    time.Time      `db:"submit_at"`
	BuilderName string         `db:"builder_name"`
}
