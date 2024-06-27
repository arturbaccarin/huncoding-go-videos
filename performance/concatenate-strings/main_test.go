package main

import (
	"fmt"
	"strings"
	"testing"
)

const (
	smallString = "ConcatenateStrings"
	longString  = "ConcatenateStringsConcatenateStringsConcatenateStringsConcatenateStringsConcatenateStrings"
)

func generateStringArray(s string) (data []string) {
	for i := 0; i < 100; i++ {
		data = append(data, s)
	}

	return
}

func BenchmarkWithSprintfSmallString(b *testing.B) {
	data := generateStringArray(smallString)

	var s string

	for n := 0; n < b.N; n++ {
		s = fmt.Sprintf(s, data)
	}
}

func BenchmarkWithSprintfLongString(b *testing.B) {
	data := generateStringArray(longString)

	var s string

	for n := 0; n < b.N; n++ {
		s = fmt.Sprintf(s, data)
	}
}

func BenchmarkWithOperatorSmallString(b *testing.B) {
	data := generateStringArray(smallString)

	f := func(s []string) (allStr string) {
		for _, x := range s {
			allStr += x
		}

		return
	}

	for n := 0; n < b.N; n++ {
		_ = f(data)
	}
}

func BenchmarkWithOperatorLongString(b *testing.B) {
	data := generateStringArray(longString)

	f := func(s []string) (allStr string) {
		for _, x := range s {
			allStr += x
		}

		return
	}

	for n := 0; n < b.N; n++ {
		_ = f(data)
	}
}

func BenchmarkWithJoinSmallString(b *testing.B) {
	data := generateStringArray(smallString)

	var s string

	for n := 0; n < b.N; n++ {
		s = strings.Join(data, "")
		_ = s
	}
}

func BenchmarkWithJoinLongString(b *testing.B) {
	data := generateStringArray(longString)

	var s string

	for n := 0; n < b.N; n++ {
		s = strings.Join(data, "")
		_ = s
	}
}

func BenchmarkWithBuilderSmallString(b *testing.B) {
	data := generateStringArray(smallString)
	var sb strings.Builder
	var s string

	for n := 0; n < b.N; n++ {
		for _, a := range data {
			sb.WriteString(a)
		}
		s = sb.String()
		_ = s
	}
}

func BenchmarkWithBuilderLongString(b *testing.B) {
	data := generateStringArray(longString)
	var sb strings.Builder
	var s string

	for n := 0; n < b.N; n++ {
		for _, a := range data {
			sb.WriteString(a)
		}
		s = sb.String()
		_ = s
	}
}
