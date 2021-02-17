package main

import (
	"fmt"
)

// Variáveis, Valores & Tipos: O pacote fmt

func main() {

	x := "oi"
	y := "bom dia"

	z := fmt.Sprint(x, " ", y)

	fmt.Println(z)
}
