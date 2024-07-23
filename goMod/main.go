package main

import (
	"fmt"

	"github.com/aryelzx/learning-GO/matematica"
)

func main() {
	s := matematica.Soma(10, 20)
	fmt.Printf("Resultado: %v\n", s)
	fmt.Println(matematica.A)

	carro := matematica.Carro{
		Marca: "Fiat",
	}

	fmt.Println(carro.Marca)
	fmt.Println(carro.Andar())

}
