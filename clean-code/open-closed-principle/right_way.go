package openclosedprinciple

type FuncPJ struct{}
type FuncCLT struct{}
type FuncEstagiario struct{}

type Funcionario interface {
	CalcularSalario() float64
}

func (f *FuncPJ) CalcularSalario() float64 {
	return 0.0
}

func (f *FuncCLT) CalcularSalario() float64 {
	return 0.0
}

func (f *FuncEstagiario) CalcularSalario() float64 {
	return 0.0
}

func FolhaDePagamento2(funcionario Funcionario) float64 {
	return funcionario.CalcularSalario()
}

func main() {
	funcionario := &FuncPJ{}
	funcionario.CalcularSalario()

	FolhaDePagamento2(funcionario)
}
