package interfacesegregationprinciple

type TrabalhoTasks interface {
	Entrar()

	ComecarATrabalhar()

	ContinuarATrabalhar()

	Sair()
}

type HumanoTasks interface {
	TrabalhoTasks

	PausaProCafe()

	Almocar()
}

type RoboTasks interface {
	TrabalhoTasks

	VerificarOleo()

	CarregarBateria()
}
