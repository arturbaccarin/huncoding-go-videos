package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	load := func() {
		fmt.Println("executando codigo init")
	}

	// var mu sync.Mutex
	// var passed bool

	var once sync.Once

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			once.Do(load)
			/*
				mu.Lock()
				if !passed {
					load()
					passed = true
				}
				mu.Unlock()
			*/
		}()
	}

	wg.Wait()
}
