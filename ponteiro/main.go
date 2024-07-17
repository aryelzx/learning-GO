package main

import "fmt"

type Conta struct {
	saldo int
}

func (c *Conta) simular(valor int) int {
	c.saldo += valor
	return c.saldo
}

func NewConta() *Conta {
	return &Conta{saldo: 0}
}

func soma(a, b *int) int {
	*a = 50 //atribui novo valor da memória
	return *a + *b
}

func Memory() {
	minhaVar1 := 10
	minhaVar2 := 20
	println(minhaVar1)                    //return 10
	println(soma(&minhaVar1, &minhaVar2)) //return 50 + 20 = 70
}

func main() {
	// Memória -> Endereço -> Valor

	//variavel -> ponteiro que tem um endereço na memória => valor
	a := 10
	var pointer *int = &a
	*pointer = 20
	b := &a
	fmt.Println(b, "=>", *b)
	fmt.Println(pointer, "=>", *pointer)

	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	Memory()

	conta := Conta{
		saldo: 100,
	}
	conta.simular(200)
	println(conta.saldo)
}
