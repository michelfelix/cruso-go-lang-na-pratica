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

//Naked return - Como pode ser observado, basicamente naked return não especifica o que vai ser o retorno porque fica subentendido devido ao atribuir o tipo e valor ao parâmetro de area e perímetro
// func calculateRectangleParams(width, height float64) (area, perimeter float64) {
// 	area = width * height
// 	perimeter = 2 * (width + height)
// 	return
// }

//Função anônima
//Função em Go é um valor de primeira classe, como uma string, int, etc. Isso significa que você pode atribuir uma função a uma variável, passá-la como argumento para outra função e retorná-la de outra função.
//Uma função anônima é uma função sem nome, que pode ser definida e chamada em tempo de execução.
var mult = func(a, b int) int {
	return a * b
}

func operate(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}

func sum (a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func getMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

//Parâmetros variádicos
func sumAll(nums ...int) int {
	sum := 0

	for _, num := range nums {
		sum += num
	}

	return sum
}

func printWithPrefix(prefix string, msgs ...string) {
	for _, msg := range msgs {
		fmt.Println(prefix, msg)
	}
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

	//Início impressão função anônima
	
	var double = getMultiplier(2)

	fmt.Println("Resultado da multiplicação: ", mult(10, 5))
	fmt.Println("Resultado da soma: ", operate(10, 5, sum))
	fmt.Println("Resultado da subtração: ", operate(10, 5, sub))
	fmt.Println("Resultado da duplicação: ", double(10))

	//Fim impressão função anônima

	//Início impressão valores variádicos
	
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(sumAll(nums...))
	printWithPrefix("Michel diz: ", "Oi", "Eu sou o Michel", "Prazer em conhecer!")
	//Fim impressão valores variádicos
}