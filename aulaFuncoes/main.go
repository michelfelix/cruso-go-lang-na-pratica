package main

import "fmt"

//Função de retorno de múltiplos valores
func isDivisibleBy2(num1, num2 float64) (float64, error) {
	if num2 == 0 {
		return 0, fmt.Errorf("Não é possível dividir por zero")
	}
	return num1 / num2, nil
}

//Named return 
func calculateRectngleParams(width, height float64) (area, perimeter float64) {
	area = width * height
	perimeter = 2* (width + height)

	return area, perimeter
}

func main() {
	//Início impressão função multiplos retornos
	var result1, err1 = isDivisibleBy2(10, 7)
	if err1 != nil {
		fmt.Println(err1)
	}else {
		fmt.Println(result1)
	}
	
	//Fim impressão função multiplos retornos

	//Início impressão named return
	
	var area, perimeter = calculateRectngleParams(10, 5)
	fmt.Println("Área do retângulo: ", area)
	fmt.Println("Perímetro do retângulo: ", perimeter)
	
	//Fim impressão named return
}