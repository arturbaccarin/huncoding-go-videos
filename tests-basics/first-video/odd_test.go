package firstvideo

import "testing"

// go test ./... -cover
func TestIsOdd_return_odd(t *testing.T) {

	// definindo variáveis (given)
	var value int64 = 3

	// chamar as funções necessarios (when)
	numberType := isOdd(value)

	// verificar o resultado (then)
	if numberType != ODD_KEYWORD {
		t.Errorf("%s - Expected %s, got %s", t.Name(), ODD_KEYWORD, numberType)
		return
	}
}

func TestIsOdd_return_even(t *testing.T) {

	// definindo variáveis (given)
	var value int64 = 4

	// chamar as funções necessarios (when)
	numberType := isOdd(value)

	// verificar o resultado (then)
	if numberType != EVEN_KEYWORD {
		t.Errorf("%s - Expected %s, got %s", t.Name(), ODD_KEYWORD, numberType)
		return
	}
}
