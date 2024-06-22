package factory

import "fmt"

type Veiculo interface {
	PrintarQuantidaDeRodas()
}

type Carro struct {
	qtdRodas int
}

func (c *Carro) PrintarQuantidaDeRodas() {
	fmt.Printf("Carro com %d rodas\n", c.qtdRodas)
}

type Moto struct {
	qtdRodas int
}

func (c *Moto) PrintarQuantidaDeRodas() {
	fmt.Printf("Moto com %d rodas\n", c.qtdRodas)
}

func NewVeiculo(qtdRodas int) Veiculo {
	if qtdRodas == 2 {
		return &Moto{qtdRodas}
	}

	if qtdRodas == 4 {
		return &Carro{qtdRodas}
	}

	fmt.Printf("Quantidade de rodas inválida: %d\n", qtdRodas)
	return nil
}

func main2() {
	v := NewVeiculo(4)
	v.PrintarQuantidaDeRodas()
}
