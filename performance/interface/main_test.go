package main

import "testing"

var (
	values = [6]string{"one", "two", "3", "4", "five", "six"}
)

func BenchmarkDefineMapFirstWay(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DefineMapFirstWay(values)
	}
}

func BenchmarkDefineMapSecWay(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DefineMapSecWay(values)
	}
}

func BenchmarkDefineMapThirdWay(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DefineMapThirdWay(values)
	}
}
