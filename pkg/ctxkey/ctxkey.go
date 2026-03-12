package ctxkey

import "context"

func Get[V any](ctx context.Context, key any) V {
	v := ctx.Value(key)
	if v == nil {
		var defaultVal V
		return defaultVal
	}

	casted, _ := v.(V)

	return casted
}

func GetCheck[V any](ctx context.Context, key any) (V, bool) {
	v := ctx.Value(key)
	if v == nil {
		var defaultVal V
		return defaultVal, false
	}

	casted, ok := v.(V)

	return casted, ok
}

type Setter struct {
	ctx context.Context // nolint: containedctx
}

func NewSetter(parent context.Context) *Setter {
	return &Setter{
		ctx: parent,
	}
}

func (s *Setter) Set(k, v any) *Setter {
	s.ctx = context.WithValue(s.ctx, k, v)

	return s
}

func (s *Setter) Ctx() context.Context {
	return s.ctx
}
