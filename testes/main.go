package main

import "fmt"

func main() {
	for x := 0; x < 3; x++ { // x++ é uma instrução, nunca uma expressão
		fmt.Println("iteração", x)
	}
}
