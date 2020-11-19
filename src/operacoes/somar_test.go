package operacoes_test

import (
	"testing"

	"github.com/robsantossilva/desafio_ci_codeeducation/operacoes"
)

func TestSomar(t *testing.T) {
	result := operacoes.Somar(5, 5)
	if result != 10 {
		t.Errorf("Soma inválida. Resultado esperado: %v, resultado obtido: %v", 10, result)
	}
}
