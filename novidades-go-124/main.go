// https://youtu.be/dLTpsusHkxs

package main

import "fmt"

type Lista[T any] = []T

func main() {
	var numeros Lista[int] = []int{1, 2, 3}
	var palavras Lista[string] = []string{"a", "b", "c"}

	fmt.Println(numeros)
	fmt.Println(palavras)
}
