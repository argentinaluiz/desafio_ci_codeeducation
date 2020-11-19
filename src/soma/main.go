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

//Somar - somar valores
func Somar(num1 float64, num2 float64) float64 {
	return num1 + num2
}
