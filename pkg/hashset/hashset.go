package hashset

import (
	"maps"
	"slices"
)

type HashSet[K comparable] struct {
	m map[K]struct{}
}

func New[K comparable]() *HashSet[K] {
	return &HashSet[K]{
		m: map[K]struct{}{},
	}
}

func (h *HashSet[K]) Contains(k K) bool {
	_, ok := h.m[k]

	return ok
}

func (h *HashSet[K]) Add(k K) {
	h.m[k] = struct{}{}
}

func (h *HashSet[K]) Remove(k K) {
	delete(h.m, k)
}

func (h *HashSet[K]) Size() int {
	return len(h.m)
}

func (h *HashSet[K]) Clear() {
	clear(h.m)
}

func (h *HashSet[K]) Keys() []K {
	return slices.Collect(maps.Keys(h.m))
}
