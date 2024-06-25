// Package panicrecover is used to learn about go doc
// This is a package doc

package panicrecover

import "fmt"

const (
	// TEST_CONST is a constant
	TEST_CONST = "TEST_CONST"
)

// assertValue does something
func assertValue(value1, value2 string) {
	if value1 == value2 {
		fmt.Println("são iguais")
	}

	panic("não são iguais")
}

// godoc -http=:6060 -> Command to start the documentation server
