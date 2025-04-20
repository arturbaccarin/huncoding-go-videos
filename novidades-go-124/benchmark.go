package main

import "testing"

func TestMain(t *testing.T) {

}

// old format
func BenchmarkHuncoding(b *testing.B) {
	for i := 0; i < b.N; i++ {
		primeNumbers(num)
	}
}

// new format
func BenchmarkPrimeNumbers(b *testing.B) {
	for b.Loop() {
		primeNumbers(num)
	}
}
