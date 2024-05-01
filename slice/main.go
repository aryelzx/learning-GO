package main

import "fmt"

//interpretamos `:` como um ponto de corte.

func main() {
	s := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20} //slice com 5 posições do tipo int

	//printa o slice completo e o tamanho e capacidade do slice
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	//do inicio para a direita ele transforma em um slice de 0 itens. mas não reduz a capacidade.
	fmt.Printf("len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0])

	// do inicio para a direita ele transforma em um slice de 4 itens. mas não reduz a capacidade.
	fmt.Printf("len=%d cap=%d %v\n", len(s[:4]), cap(s[:4]), s[:4])

	//pula os 2 primeiros itens e transforma em um slice de 6 itens. reduzindo a capacidade.
	fmt.Printf("len=%d cap=%d %v\n", len(s[2:]), cap(s[2:]), s[2:])

	// incrementando no slice
	s = append(s, 22, 24, 26)
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
