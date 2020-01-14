package ConcurrentMap

import (
	"fmt"
	"sync"
)

type RWMap struct {
	sync.RWMutex
	list map[interface{}]interface{}
}

func (m *RWMap) Set(k interface{}, v interface{}) {
	m.Lock()
	m.list[k] = v
	m.Unlock()
}
func (m *RWMap) Get(k interface{}) interface{} {
	m.Lock()
	defer m.Unlock()
	return m.list[k]
}

func New() *RWMap {
	return &RWMap{list: make(map[interface{}]interface{})}
}

func init() {
	m := New()
	for i := 1; i < 100; i++ {
		k := i
		go func() {
			fmt.Println("ConcurrentMapB", k, m.Get(k))
			m.Set(k, k)
			fmt.Println("ConcurrentMapA", k, m.Get(k))
		}()
	}
}
