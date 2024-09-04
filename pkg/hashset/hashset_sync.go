package hashset

import (
	"sync"
)

type SyncHashSet[K comparable] struct {
	rw sync.RWMutex
	s  *HashSet[K]
}

func NewSync[K comparable]() *SyncHashSet[K] {
	return &SyncHashSet[K]{
		rw: sync.RWMutex{},
		s:  New[K](),
	}
}

func (sh *SyncHashSet[K]) Contains(k K) bool {
	sh.rw.RLock()
	defer sh.rw.RUnlock()

	return sh.s.Contains(k)
}

func (sh *SyncHashSet[K]) ContainsOrAdd(k K) (isContains bool) {
	sh.rw.Lock()
	defer sh.rw.Unlock()
	if sh.s.Contains(k) {
		return true
	}

	sh.s.Add(k)

	return false
}

func (sh *SyncHashSet[K]) Add(k K) {
	sh.rw.Lock()
	defer sh.rw.Unlock()

	sh.s.Add(k)
}

func (sh *SyncHashSet[K]) Remove(k K) {
	sh.rw.Lock()
	defer sh.rw.Unlock()

	sh.s.Remove(k)
}

func (sh *SyncHashSet[K]) Size() int {
	sh.rw.RLock()
	defer sh.rw.RUnlock()

	return sh.s.Size()
}

func (sh *SyncHashSet[K]) Clear() {
	sh.rw.Lock()
	defer sh.rw.Unlock()

	sh.s.Clear()
}

func (sh *SyncHashSet[K]) Keys() []K {
	sh.rw.RLock()
	defer sh.rw.RUnlock()

	return sh.s.Keys()
}
