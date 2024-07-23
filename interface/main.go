package main

import "fmt"

/*
	Trata e aceita qualquer tipo. *Generics substitui interfaces vazias*
*/

func main() {
	var x interface{} = 10
	var y interface{} = "Hello, World"

	showType(x)
	showType(y)

	var minhaVar interface{} = "Aryel Ramos"
	println(minhaVar.(string)) //type assertion
	res, ok := minhaVar.(int)  //res - valor convertido para inteiro    ok - caso seja convertido com sucesso retorna true.
	fmt.Printf(" O valor de res é %v e o resultado de ok é %v", res, ok)
}

func showType(t interface{}) {
	fmt.Printf("O tipo da variavél é %T e o valor é %v\n", t, t)
}
