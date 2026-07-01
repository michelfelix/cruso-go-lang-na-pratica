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
	perimeter = 2 * (width + height)

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

func sum(a, b int) int {
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

func printWithPrefix(prefix string, msgs ...string) { //Aqui podemos observar que o valor variádico é o último parâmetro da função
	for _, msg := range msgs {
		fmt.Println(prefix, msg)
	}
}

//Closures
func makeIDGenerator() func() int {
	id := 0

	return func() int {
		id++
		return id
	}
}

//Explicação breve sobre clojure, o valor permanece enquanto a aplicação entende que o valor vai ser utilizado, quem gerencia isso é o escape analysis.
//Segue um fluxograma de como é o funcionamento desse gerenciamento entre stack, heap e escape analysis
//Você declara uma variável
//        │
//        ▼
//
//Compilador faz Escape Analysis
//        │
//   ┌────┴─────┐
//   │          │
//  Não		 Escapa
// escapa		|
//   │          │
//  Stack      Heap
//              │
//              ▼
//      Garbage Collector gerencia

//Defer
func getValue() int {
	fmt.Println("chamando getValue()")
	return 42
}
func main() {
	//Início impressão função multiplos retornos
	var result1, err1 = isDivisibleBy2(10, 7)
	if err1 != nil {
		fmt.Println(err1)
	} else {
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

	//Início impressão closures

	idGen1 := makeIDGenerator()
	idGen2 := makeIDGenerator()
	fmt.Println("ID1: ", idGen1())
	fmt.Println("ID2: ", idGen2())
	fmt.Println("Novo ID1: ", idGen1())
	fmt.Println("Novo ID2: ", idGen2())
	//Fim impressão closures

	//Início impressão defer
	defer fmt.Println("Valor obtido", getValue()) // O defer é utilizado para adiar a execução de uma função até que a função que a chamou retorne. No caso, o getValue() será chamado imediatamente e seu valor armazenado no momento do seu retorno, mas o fmt.Println só será executado no final da função main.
	fmt.Println("Função main retornando")
	//Fim impressão defer

	//Início impressão panic e recover

	defer func() {
		r := recover()

		if r != nil {
			fmt.Println("Recuperado do panic: ", r) //Este é o momento que temos para finalizar nosso programa caso o panic alerte algum problema,podendo encerrar conexões, dar alertas ou até mesmo salvar dados em um banco de dados, etc.
		}
	}()

	fmt.Println("Antes do panic")
	defer fmt.Println("Defer ainda é retornado mesmo após o panic") //Nesse exemplo de defer, percebemos que ele é "last in first out", ou seja, ele é executado na ordem inversa da sua declaração
	panic("Ocorreu um erro inesperado!")                            //O panic para a aplicação e o código abaixo dele não sera executado.
	fmt.Println("Esse código nunca será executado")

	//Fim impressão panic e recover
}
