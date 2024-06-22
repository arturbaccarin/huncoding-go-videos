package interfacesegregationprinciple

type Trabalho interface {
	Entrar()

	ComecarATrabalhar()

	PausaProCafe()

	VerificarOleo()

	Almocar()

	CarregarBateria()

	ContinuarATrabalhar()

	Sair()
}

type Robo struct{}

func (*Robo) Entrar() {}

func (*Robo) ComecarATrabalhar() {}

func (*Robo) PausaProCafe() {
	// não faz sentido
	panic("not implemented")
}

func (*Robo) VerificarOleo() {}

func (*Robo) Almocar() {
	// não faz sentido
	panic("not implemented")
}

func (*Robo) CarregarBateria() {}

func (*Robo) ContinuarATrabalhar() {}

func (*Robo) Sair() {}

type Humano struct{}

func (*Humano) Entrar() {}

func (*Humano) ComecarATrabalhar() {}

func (*Humano) PausaProCafe() {}

func (*Humano) VerificarOleo() {
	// não faz sentido
	panic("not implemented")
}

func (*Humano) Almocar() {}

func (*Humano) CarregarBateria() {
	// não faz sentido
	panic("not implemented")
}

func (*Humano) ContinuarATrabalhar() {}

func (*Humano) Sair() {}
