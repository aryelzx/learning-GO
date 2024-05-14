package main

import "fmt"

type Pessoa struct {
    Nome     string
    Idade    int
    Profissao string
}

func main() {
	//criando uma instância da struct Pessoa
	aryel := Pessoa{
		Nome: "Aryel",
		Idade: 21,
		Profissao: "Desenvolvedor",
	}

	//acessando os campos da struct
	fmt.Println("Nome:", aryel.Nome)
	fmt.Println("Idade: ", aryel.Idade)
	fmt.Println("Profissão: ", aryel.Profissao)

	aryel.Profissao = "Desenvolvedor FullStack"
	fmt.Println("Nova Profissão:", aryel.Profissao)
}
