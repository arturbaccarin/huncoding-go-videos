package main

// Não existe apenas uma forma de resolver um certo problema.

func naturalNumbers(n int) (sum int) {
	for i := 1; i <= n; i++ {
		sum += i
	}

	return
}

func naturalNumbersPerf(n int) int {
	return n * (n + 1) / 2
}
