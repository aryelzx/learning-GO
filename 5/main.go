// range é uma palavra-chave usada para iterar sobre elementos 
// em várias estruturas de dados, incluindo arrays, slices, maps e channels.
// Ele fornece uma maneira conveniente de percorrer os elementos de uma 
// coleção sem a necessidade de manipulação explícita de índices.

package main

import "fmt"

type ID int
	
var (
	a string  //por padrão é vazio " "
	b bool    //por padrão ele é false
	c int     //por padrão ele é 0
	d float64 //por padrão ele é 0.0
	f ID      = 1
)

func main() {
	var meuArray [3]int //array com 3 posições fixas do tipo int

	meuArray[0] = 1 //atribuindo valores para as posições do array
	meuArray[1] = 2
	meuArray[2] = 3
	
	for i, v := range meuArray { 
		fmt.Printf("O valor do índice é %d e o valor é %d\n", i, v)
		//%d é um marcador de posição para um número inteiro
		//%n é um marcador de posição para um número
	}
}
