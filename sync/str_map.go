package sync

import (
	"sync"
)

type AsyncStrMap struct {
	m map[string]interface{}
	x sync.RWMutex
}

func NewAsyncStrMap() AsyncStrMap {
	return AsyncStrMap{
		m: map[string]interface{}{},
	}
}

func (a *AsyncStrMap) Set(key string, value interface{}) bool {
	defer a.x.Unlock()
	a.x.Lock()
	if _, ok := a.m[key]; ok {
		return false
	}
	a.m[key] = value
	return true
}

func (a *AsyncStrMap) Get(key string) (interface{}, bool) {
	defer a.x.RUnlock()
	a.x.RLock()

	v, ok := a.m[key]

	return v, ok
}

func (a *AsyncStrMap) Len() int {
	defer a.x.RUnlock()
	a.x.RLock()

	return len(a.m)
}
