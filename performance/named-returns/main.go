package main

import (
	"fmt"
	"time"
)

func declareReturnTypeOnly() string {
	execTime := time.Now()

	var allStr string

	s := [5]string{"one", "two", "3", "4", "five"}
	for _, str := range s {
		allStr += str
	}

	fmt.Printf("declareReturnTypeOnly took %6s\n", time.Since(execTime))

	return allStr
}

func declareReturnNameTypeWithLog() (allStr string) {
	execTime := time.Now()

	s := [5]string{"one", "two", "3", "4", "five"}
	for _, str := range s {
		allStr += str
	}

	fmt.Printf("declareReturnNameTypeWithLog took %6s\n", time.Since(execTime))

	return
}

func main() {
	declareReturnTypeOnly()
	declareReturnNameTypeWithLog()
}
