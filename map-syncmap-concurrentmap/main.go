// https://youtu.be/yFJqQkrAj9w
package main

import (
	"fmt"
	"sync"
)

func standardMap() {
	m := make(map[int]int)
	var wg sync.WaitGroup

	for i := range 100 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			m[i] = i * 10
		}(i)
	}

	wg.Wait()
	fmt.Println("Standard Map:", m)
}

func syncMap() {
	var sm sync.Map
	var wg sync.WaitGroup

	for i := range 10 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			sm.Store(i, i*10)
		}(i)
	}

	wg.Wait()

	sm.Range(func(key, value interface{}) bool {
		fmt.Println("Sync Map:", key, value)
		return true
	})

	sm.CompareAndSwap(1, 10, 100)
}

type ConcurrentMap struct {
	mu sync.RWMutex
	m  map[int]int
}

func (c *ConcurrentMap) Store(key, value int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.m[key] = value
}

func (c *ConcurrentMap) Load(key int) (int, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, ok := c.m[key]
	return value, ok
}

func concurrentMap() {
	cm := &ConcurrentMap{
		m: make(map[int]int),
	}

	var wg sync.WaitGroup
	for i := range 10 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			cm.Store(i, i*10)
		}(i)
	}

	wg.Wait()

	for i := range 10 {
		value, ok := cm.Load(i)
		if ok {
			fmt.Println("Concurrent Map:", i, value)
		} else {
			fmt.Println("Concurrent Map: Key", i, "not found")
		}
	}
}

func main() {

}
