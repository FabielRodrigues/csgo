package main

import (
	"fmt"
)

var y = "Olá bom dia"

func main() {
	x := 10

	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n\n", y, y)

	x = 20
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n\n", y, y)
}
