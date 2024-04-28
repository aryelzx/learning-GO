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
	fmt.Printf("O tipo E é de %v, - %T ", f, f)
}
