package main

import (
	"testing"
)

func TestSomar(t *testing.T) {
	result := somar(5, 5)

	if result != 11 {
		t.Errorf("Soma inválida. Resultado esperado: %v, resultado obtido: %v", 10, result)
	}
}
