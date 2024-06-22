package openclosedprinciple

import "reflect"

type FuncionarioCLT struct{}

type Estagiario struct{}

// e se precisar adicionar o FuncionarioPJ?

func FolhaDePagamento[TipoFuncionario FuncionarioCLT | Estagiario](funcionario TipoFuncionario) float64 {

	if reflect.TypeOf(FuncionarioCLT{}) == reflect.TypeOf(funcionario) {

		// calcular salario + beneficios
		return 0.0
	}

	// calcular bolsa auxilio de estagiario
	return 0.0
}
