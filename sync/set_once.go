package sync

import (
	"sync"
	"sync/atomic"
)

// A SetOnce is an object that will store a value once, any attempts to change
// the stored value will fail.
type SetOnce struct {
	v   interface{}
	m   sync.Mutex
	set uint32
}

// Attempt to set a value. If this is the first time the vlaue has been set,
// then it will store the value and return true. Any other subsequent tries will
// ignore the values and return false.
func (o *SetOnce) Set(v interface{}) bool {
	if atomic.LoadUint32(&o.set) == 1 {
		return false
	}

	o.m.Lock()
	defer o.m.Unlock()
	if o.set == 0 {
		defer atomic.StoreUint32(&o.set, 1)
		o.v = v

		return true
	}

	return false
}

// Query the stored value. If no value has been set then v will be nil and ok
// will be false.
func (o *SetOnce) Get() (v interface{}, ok bool) {
	if atomic.LoadUint32(&o.set) == 0 {
		return nil, false
	}

	return o.v, true
}
