package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

type Recurso struct {
	nome string
}

func liberarRecurso(r *Recurso) {
	fmt.Printf("Recurso %s liberado\n", r.nome)
}

func main() {
	r := &Recurso{nome: "Banco de dados"}

	runtime.SetFinalizer(r, liberarRecurso)

	fmt.Println("Recurso criado, aguardando GC...")

	r = nil

	runtime.GC()

	time.Sleep(2 * time.Second)
}

// new
func openFile(filename string) *os.File {
	f, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	runtime.AddCleanup(f, func() {
		fmt.Println("Arquivo fechado")
		f.Close()
	})

	return f
}

func main() {
	file := openFile("arquivo.txt")

	if file != nil {
		fmt.Println("Arquivo aberto")
	}

	file = nil
	runtime.GC()
}
