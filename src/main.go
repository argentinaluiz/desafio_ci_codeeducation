package main

import (
	"fmt"

	"github.com/robsantossilva/desafio_ci_codeeducation/src/operacoes"
)

func main() {
	result := operacoes.Somar(5, 5)
	s := fmt.Sprintf("%.2f", result)
	fmt.Println(s)
}
