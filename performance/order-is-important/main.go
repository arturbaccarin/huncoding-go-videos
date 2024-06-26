package main

import (
	"fmt"
	"time"
	"unsafe"
)

type commonStruct struct {
	createdAt time.Time
	updatedAt time.Time
	timeout   time.Duration
	jsonStr   []byte
}

type Employee1 struct {
	IsActive  bool
	Age       int64
	IsMarried bool
	Name      string
	weight    int32
	height    int16
	PhotoLen  float64
	PhotoWid  float32
	intNum    int
	length    int8
	common    commonStruct
}

type Employee2 struct {
	Name      string
	Age       int64
	intNum    int
	PhotoLen  float64
	PhotoWid  float32
	weight    int32
	height    int16
	length    int8
	IsMarried bool
	IsActive  bool
	common    commonStruct
}

var e1 Employee1
var e2 Employee2

func main() {
	fmt.Printf("Size of %T struct: %d bytes\n", e1, unsafe.Sizeof(e1))
	fmt.Printf("Size of %T struct: %d bytes\n", e2, unsafe.Sizeof(e2))
}
