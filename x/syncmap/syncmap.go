package syncmap

import "sync"

type SyncMap[K comparable, V any] struct {
	data map[K]V
	l    sync.RWMutex
}

func New[K comparable, V any]() SyncMap[K, V] {
	return SyncMap[K, V]{
		data: make(map[K]V),
	}
}

func (m *SyncMap[K, V]) Store(k K, v V) {
	m.l.Lock()
	defer m.l.Unlock()
	m.data[k] = v
}

func (m *SyncMap[K, V]) Delete(k K) {
	m.l.Lock()
	defer m.l.Unlock()
	delete(m.data, k)
}

func (m *SyncMap[K, V]) Update(k K, fn func(V) (V, error)) (bool, error) {
	m.l.Lock()
	defer m.l.Unlock()

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
	m.l.RLock()
	defer m.l.RUnlock()
	v, ok = m.data[k]
	return v, ok
}
