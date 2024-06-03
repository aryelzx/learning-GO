package main

import "fmt"

func main() {
	// Memória -> Endereço -> Valor

	//variavel -> ponteiro que tem um endereço na memória => valor
	a := 10
	var pointer *int = &a
	*pointer = 20
	b := &a
	fmt.Println(b)
	fmt.Println(pointer)

	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
