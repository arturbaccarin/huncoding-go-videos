package main

// Não existe apenas uma forma de resolver um certo problema.

// Maneira 1
func naturalNumbers(n int) (sum int) {
	for i := 1; i <= n; i++ {
		sum += i
	}

	return
}

// Maneira 2 mais eficiente
func naturalNumbersPerf(n int) int {
	return n * (n + 1) / 2
}
