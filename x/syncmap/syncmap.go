package syncmap

import "sync"

type SyncMap[K comparable, V any] struct {
	data map[K]V
	rw   sync.RWMutex
}

func New[K comparable, V any]() SyncMap[K, V] {
	return SyncMap[K, V]{
		data: make(map[K]V),
	}
}

func (m *SyncMap[K, V]) Store(k K, v V) {
	m.rw.Lock()
	defer m.rw.Unlock()
	m.data[k] = v
}

func (m *SyncMap[K, V]) Delete(k K) {
	m.rw.Lock()
	defer m.rw.Unlock()
	delete(m.data, k)
}

func (m *SyncMap[K, V]) Update(k K, fn func(V) (V, error)) (bool, error) {
	m.rw.Lock()
	defer m.rw.Unlock()

	v, ok := m.data[k]
	if !ok {
		return false, nil
	}

	appliedV, err := fn(v)
	if err != nil {
		return false, err
	}

	m.data[k] = appliedV

	return true, nil
}

func (m *SyncMap[K, V]) Load(k K) (v V, ok bool) {
	m.rw.RLock()
	defer m.rw.RUnlock()
	v, ok = m.data[k]
	return v, ok
}

func (m *SyncMap[K, V]) Keys() []K {
	m.rw.RLock()
	defer m.rw.RUnlock()

	keys := make([]K, 0, len(m.data))

	for k := range m.data {
		keys = append(keys, k)
	}

	return keys
}

func (m *SyncMap[K, V]) RangeMut(fn func(k K, v V) bool) {
	m.rw.Lock()
	defer m.rw.Unlock()

	for k, v := range m.data {
		if !fn(k, v) {
			break
		}
	}
}
