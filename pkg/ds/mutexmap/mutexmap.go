package mutexmap

import "sync"

type MutexMap[K comparable, V any] struct {
	data map[K]V
	l    sync.RWMutex
}

func New[K comparable, V any]() MutexMap[K, V] {
	return MutexMap[K, V]{
		data: make(map[K]V),
	}
}

func (m *MutexMap[K, V]) Store(k K, v V) {
	m.l.Lock()
	defer m.l.Unlock()
	m.data[k] = v
}

func (m *MutexMap[K, V]) Delete(k K) {
	m.l.Lock()
	defer m.l.Unlock()
	delete(m.data, k)
}

func (m *MutexMap[K, V]) Load(k K) (v V, ok bool) {
	m.l.RLock()
	defer m.l.RUnlock()
	v, ok = m.data[k]
	return v, ok
}
