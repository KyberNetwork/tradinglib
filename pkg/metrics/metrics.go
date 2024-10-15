package metrics

import (
	"context"

	kybermetric "github.com/KyberNetwork/kyber-trace-go/pkg/metric"
	"go.opentelemetry.io/otel/metric"
)

var m metric.Meter

func init() {
	m = kybermetric.Meter()
}

func RecordFloat64Histogram(ctx context.Context, name string, value float64) error {
	hist, err := m.Float64Histogram(name)
	if err != nil {
		return err
	}
	hist.Record(ctx, value)
	return nil
}

func RecordFloat64Gause(ctx context.Context, name string, value float64) error {
	_, err := m.Float64ObservableGauge(name, metric.WithFloat64Callback(func(ctx context.Context, fo metric.Float64Observer) error {
		fo.Observe(value)
		return nil
	}))
	return err
}

func RecordCounter(ctx context.Context, name string, value int64) error {
	counter, err := m.Int64Counter(name)
	if err != nil {
		return err
	}
	counter.Add(ctx, value)
	return nil
}

func RecordUpdownCounter(ctx context.Context, name string, value int64) error {
	counter, err := m.Int64UpDownCounter(name)
	if err != nil {
		return err
	}
	counter.Add(ctx, value)
	return nil
}
