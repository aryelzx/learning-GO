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

type Empresa struct {
	Nome string
}
type Funcionario interface {
	Contratar()
	Desativar()
}

func (f Empresa) Contratar() {
	funcionarios := make(map[string]string)
	funcionarios["nome"] = "Aryel"
	funcionarios["cargo"] = "Desenvolvedor"
	fmt.Println("Funcionário contratado")
	fmt.Printf("Nome: %s\nCargo: %s\n Empresa: %s", funcionarios["nome"], funcionarios["cargo"], f.Nome)
}

func Contratacao(empresa Empresa) {
	empresa.Contratar()
}

func (e Empresa) Desativar() { //declaro o metodo para a struct Empresa
	e.Nome = "Normaltech"
	fmt.Printf("Empresa %s desativada", e.Nome)
}

func Desativacao(empresa Empresa) { // function que recebe um parametro do tipo Empresa e chama o método Desativar
	empresa.Desativar()
}

func (p Pessoa) Desativar() { //método
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

	//acessando o método da struct Empresa
	minhaEmpresa := Empresa{}
	Desativacao(minhaEmpresa)

	//acessando o método da interface Funcionario
	contrato := Empresa{
		Nome: "NetsoftWay",
	}
	Contratacao(contrato)

}
