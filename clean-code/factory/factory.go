package factory

import "fmt"

type carro struct {
	qtdRodas int
	preco    float64
}

func (c *carro) SetPreco(preco float64) {
	if preco < 5000 || preco > 1000000 {
		fmt.Println("Preço inválido")
		return
	}

	c.preco = preco
}

func (c *carro) SetQtdRodas(qtdRodas int) {
	if qtdRodas < 0 || qtdRodas > 4 {
		fmt.Println("Quantidade de rodas invalida")
		return
	}

	c.qtdRodas = qtdRodas
}

func NewCar() carro {
	return carro{4, 150.000}
}

// se for em outro pacote
func main() {
	c := NewCar()
	c.SetPreco(123456)
}
