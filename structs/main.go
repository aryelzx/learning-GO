package main

import "fmt"

type Endereco struct {
	Logradouro string
	Bairro     string
	Cidade     string
	Estado     string
}

type Pessoa struct {
	Nome      string
	Idade     int
	Profissao string
	Ativo     bool
	Adress    Endereco
	Cliente
}

type Cliente struct {
	Convenio string
	Planos   []string
	Endereco
}

func (p Pessoa) Desativar() {
	p.Ativo = false
	fmt.Printf("O cliente %s foi desativado, %v", p.Nome, p.Ativo)
}

func main() {
	//criando uma instância da struct Pessoa
	aryel := Pessoa{
		Nome:      "Aryel",
		Idade:     21,
		Profissao: "Desenvolvedor",
		Ativo:     true,
	}

	aryel.Adress.Cidade = "Fortaleza"

	//acessando os campos da struct
	fmt.Println("Nome:", aryel.Nome)
	fmt.Println("Idade: ", aryel.Idade)
	fmt.Println("Profissão: ", aryel.Profissao)
	fmt.Println("Ativo: ", aryel.Ativo)

	aryel.Profissao = "Desenvolvedor FullStack"
	fmt.Println("Nova Profissão:", aryel.Profissao)
	//acessando o campo da struct Endereco
	fmt.Println("Cidade:", aryel.Adress.Cidade)

	//acessando o método da struct cliente
	aryel.Desativar()

}
