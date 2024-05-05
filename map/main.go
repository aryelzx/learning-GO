package main

import "fmt"

func main() {

	salarios := map[string]int{"José": 3000, "Maria": 500, "Pedro": 2000}

	delete(salarios, "José") //deleta um elemento do map
	salarios["Maria"] = 1000 //atualiza um elemento do map
	salarios["João"] = 1500  //adiciona um elemento ao map

	sal := make(map[string]int) //cria um map vazio
	sal2 := map[string]int{}    //cria um map vazio

	sal["Aryel"] = 5000  //adiciona um elemento ao map
	sal2["Aryel"] = 6000 //adiciona um elemento do map

	for chave, valor := range salarios {
		fmt.Println("Chave:", chave, "Valor:", valor)
	}

	for _, valor := range salarios { //ignora a chave com blank identifier
		fmt.Println("Valor:", valor)
	}
}
