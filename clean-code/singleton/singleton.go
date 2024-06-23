package main

import (
	"fmt"
	"sync"
)

var (
	CAPITALS = []string{"BRASILIA", "SALVADOR", "PORTO ALEGRE", "CURITIBA"}
)

type singletonDatabase struct {
	capitals []string
}

func (sd *singletonDatabase) GetCapitals() []string {
	return sd.capitals
}

/**
// outra forma, porém não permite fazer um New
func init() {

}
*/

var once sync.Once
var instance *singletonDatabase

func NewSingletonDatabase() *singletonDatabase {
	capitals := CAPITALS
	once.Do(func() {
		db := singletonDatabase{capitals}
		fmt.Println("Iniciou o objetos")
		instance = &db
	})
	fmt.Println(&instance)
	return instance
}

func main() {
	fmt.Println(NewSingletonDatabase(), NewSingletonDatabase(), NewSingletonDatabase())
}
