package main

import "fmt"

func modifyArray(arr [3]int) {
	arr[0] = 9
	fmt.Println(arr)
}

func modifySlice(slc []int) {
	slc[0] = 9
	fmt.Println(slc)
}

func main() {
	//Maneiras de declarar arrays
	var arr [3]int             //Aqui vemos o "[3]" que indica que tem um array de tamanho 3, nem mais, nem menos e não pode ser alterado(obs: o tamanho do array é parte do seu tipo)
	arr2 := [3]int{}           //Aqui vemos o "{}" que indica que tem um slice
	var arr3 = [3]int{1, 2, 3} //Aqui já declaramos um array de tamanho 3 e já inicializamos com valores
	var (
		arr4 = [3]int{}
		arr5 = [4]int{}
		// Se comparar esses arrays o restultado será false, pois o tamanho do array é parte do seu tipo, ou seja, o tamanho do array é diferente, logo são tipos diferentes
	)

	//Não é possível apagar elementos do array mas é possível alterar, dependendo do caso, podemos alterar o valor para 0 que é zero value do tipo int
	//Posição no array também é diferente, estou acostumado a iniciar com índice 0 e contar a partir dele, mas, num array de tamanho 4, para acessar o último elemento, o índice é 3,
	// pois o índice começa em 0, então, para acessar o último elemento do array, é necessário subtrair 1 do tamanho do array. Vide exemplo abaixo:
	var arr6 = [4]int{1, 2, 3, 4}
	arr6[3] = 9 //Aqui estou acessando o último elemento do array, que é o índice 3, e atribuindo o valor 9 a ele
	fmt.Println("<-- Início de declarações de arrays -->")
	fmt.Println(arr)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Printf("%T\n", arr4) //Essa impressão retorna o tipo [3]int, que é o tipo e tamanho do array, ou seja, o tamanho do array é parte do seu tipo
	fmt.Printf("%T\n", arr5) //Essa impressão retorna o tipo [4]int, que é o tipo e tamanho do array, ou seja, o tamanho do array é parte do seu tipo
	fmt.Println(arr6)
	fmt.Println("<-- Fim de declarações de arrays -->")
	//Início matriz
	var matrix [3][3]int

	matrix[0][0] = 1
	matrix[0][1] = 2
	matrix[0][2] = 3

	matrix[1][0] = 4
	matrix[1][1] = 5
	matrix[1][2] = 6

	matrix[2][0] = 7
	matrix[2][1] = 8
	matrix[2][2] = 9

	fmt.Println("<-- Início de matrizes -->")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}
	fmt.Println("<-- Fim de matrizes -->")
	//Fim matriz

	//Início passagem de array como valor

	foo1 := [3]int{1, 2, 3}

	fmt.Println("<-- Início de passagem de array como valor -->")

	fmt.Println(foo1)

	modifyArray(foo1) //Aqui estamos passando o array como valor, ou seja, uma cópia do array é passada para a função, então, qualquer alteração feita na função não afetará o array original

	fmt.Println(foo1) //Aqui vemos que o array original não foi alterado, pois a função modifyArray recebeu uma cópia do array foo, então, qualquer alteração feita na função não afetará o array original

	fmt.Println("<-- Fim de passagem de array como valor -->")
	//Fim passagem de array como valor

	//Início passagem de array como referência

	foo2 := []int{1, 2, 3}

	fmt.Println("<-- Início de passagem de array como referência -->")

	fmt.Println(foo2)

	modifySlice(foo2) //Aqui estamos passando o array como referência, ou seja, a referência do array é passada para a função, então, qualquer alteração feita na função afetará o array original

	fmt.Println(foo2) //Aqui vemos que o array original foi alterado, pois a função modifySlice recebeu a referência do array foo, então, qualquer alteração feita na função afetará o array original

	fmt.Println("<-- Fim de passagem de array como referência -->")
	//Fim passagem de array como referência
}
