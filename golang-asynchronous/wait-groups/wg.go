package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	// wg.Add(3) -> ele só vai criar 3 e esperar 3 e encerra o programa
	// wg.Add(10) -> se colocar mais, ele dá deadlock
	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func() {
			DemorarParaExecutar(2 * time.Second)
			defer wg.Done()
		}()
	}

	wg.Wait()
}

func DemorarParaExecutar(t time.Duration) {
	fmt.Println("Começando a execução do método")

	time.Sleep(t)

	fmt.Println("Finalizando a execução do método")
}
