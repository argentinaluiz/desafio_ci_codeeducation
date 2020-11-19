package main

import (
	"fmt"
)

//Somar - somar valores
func somar(num1 float64, num2 float64) float64 {
	return num1 + num2
}

func main() {
	result := somar(5, 5)
	s := fmt.Sprintf("%.2f", result)
	fmt.Println(s)
}
