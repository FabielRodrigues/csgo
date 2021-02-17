package main

import "fmt"

// Criando seu próprio tipo

type hotdog int

var b hotdog = 10

func main() {
	x := 10
	fmt.Printf("%v\n", x)
	fmt.Printf("%v", b)
}
