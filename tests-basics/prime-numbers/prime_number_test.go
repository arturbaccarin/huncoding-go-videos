package main

import (
	"fmt"
	"testing"
)

var num = 100

var table = []struct {
	input int
}{
	{input: 100},
	{input: 1000},
	{input: 74382},
	{input: 382399},
}

// go test -bench .
// para filtrar por nome: go test -bench Improved .
// para mostrar a quantidade de alocação de memória por operação: go test -bench . -benchmem
// para escolher a quantidade de vezes que você quer executar: go test -bench . -benchtime=1000x / 10s/m/h
// escutar o benchmark completo com N vezes por X vezes, sendo X: go test -bench . -count=5

func BenchmarkPrimeNumbersCoding(b *testing.B) {
	// mock
	for i := 0; i < b.N; i++ {
		primeNumbers(num)
	}
}

func BenchmarkPrimeNumbersImproved(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				isPrimeImproved(v.input)
			}
		})
	}
}
