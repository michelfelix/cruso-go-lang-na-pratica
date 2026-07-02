package main

import "fmt"

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

	fmt.Println(arr)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Printf("%T\n", arr4) //Essa impressão retorna o tipo [3]int, que é o tipo e tamanho do array, ou seja, o tamanho do array é parte do seu tipo
	fmt.Printf("%T\n", arr5) //Essa impressão retorna o tipo [4]int, que é o tipo e tamanho do array, ou seja, o tamanho do array é parte do seu tipo
	fmt.Println(arr6)
}
