// https://youtu.be/yFJqQkrAj9w
package main

import (
	"fmt"
	"sync"
)

func standardMap() {
	m := make(map[int]int)
	var wg sync.WaitGroup

	for i := range 10 {
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

}

func concurrentMap() {

}

func main() {

}
