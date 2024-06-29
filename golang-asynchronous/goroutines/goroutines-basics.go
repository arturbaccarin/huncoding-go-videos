package main

import (
	"fmt"
	"time"
)

func fun(value string) {
	for i := 0; i < 3; i++ {
		fmt.Println(value)
		time.Sleep(1 * time.Millisecond)
	}
}

func mainOld() {
	// direct call
	fun("direct call")

	// goroutine with differents variants for function call
	fgx := fun
	go fgx("goroutine - 2")

	// goroutine function call
	go fun("goroutine - 1")

	// goroutines with anonymous value call
	go func(t string) {
		fun(t)
	}("goroutine 3")

	time.Sleep(5 * time.Millisecond)

	// done
	fmt.Println("Done...")
}
