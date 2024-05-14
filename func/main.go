package main
import "fmt"

func main() {
valor, err := sum3(20, 10)
	if err != nil { //se encontra erro, printa o erro
		fmt.Println(err)
		return
	}

	fmt.Println(valor) //se não encontra erro, printa o valor

	fmt.Println(sum4(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}
	
func sum3(a, b int) (int, error) {
	if a+b >= 50 {
		return 0, fmt.Errorf("O valor %d é muito grande", a+b)
	}
	return a + b, nil
}

// ===================== funções variádicas ============================
func sum4(numeros ...int) int {
	total := 0 
	for _, numero := range numeros {
		total += numero
	}
	return total
}

