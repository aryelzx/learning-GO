package main

import "fmt"

{/*
	Trata e aceita qualquer tipo. *Generics substitui interfaces vazias*	
*/}

func main() {
	var x interface{} = 10
	var y interface{} = "Hello, World"

	showType(x)
	showType(y)
}

func showType(t interface{}) {
	fmt.Printf("O tipo da variavél é %T e o valor é %v\n", t, t)
}
