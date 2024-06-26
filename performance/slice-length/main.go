package main

import (
	"log"
	"time"
)

func declareReturnsSliceNoLength(s [5]string) (allSlice []string) {
	execTime := time.Now()

	for sum := 0; sum < 100; sum++ {
		for _, v := range s {
			allSlice = append(allSlice, v)
		}
	}

	log.Printf("declareReturnsSliceNoLength took %6s\n", time.Since(execTime))

	return
}

func declareReturnsSliceWithLength(s [5]string) (allSlice [500]string) {
	execTime := time.Now()

	for sum := 0; sum < 100; sum++ {
		for i, v := range s {
			allSlice[i] = v
		}
	}

	log.Printf("declareReturnsSliceNoLength took %6s\n", time.Since(execTime))

	return
}

func main() {
	declareReturnsSliceNoLength([5]string{"one", "two", "3", "4", "five"})
	declareReturnsSliceWithLength([5]string{"one", "two", "3", "4", "five"})
}
