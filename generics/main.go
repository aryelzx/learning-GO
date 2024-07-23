package main

//Contrancts - tipo especifico para ser substituido na tipagem

type MyNumber int

type Number interface {
	~int | ~float64 //use til to make that comparable with MyNumber type
}

func SomaInteiro[T Number](m map[string]T) T { //T - generic int or float64
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	m := map[string]int{"Aryel": 1000, "João": 2000, "Maria": 3000}
	m2 := map[string]float64{"Aryel": 10.00, "João": 20.00, "Maria": 30.00}
	m3 := map[string]MyNumber{"Aryel": 10.00, "João": 20.00, "Maria": 30.00}

	println(SomaInteiro(m))
	println(SomaInteiro(m2))
	println(SomaInteiro(m3))
	println(Compara(11, 10))
}
