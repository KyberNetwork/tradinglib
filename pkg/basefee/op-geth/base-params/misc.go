package params

// uint64ptr is a weird helper to allow 1-line constant pointer creation.
//
//go:fix inline
func uint64ptr(n uint64) *uint64 {
	return new(n)
}
