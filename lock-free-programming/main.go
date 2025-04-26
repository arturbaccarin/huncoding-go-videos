// https://youtu.be/8JOrxl0KkDs
// commit of the day :)
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup

	var counter int64

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			atomic.AddInt64(&counter, 1)
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}

/*
func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	counter := 0

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}
*/
