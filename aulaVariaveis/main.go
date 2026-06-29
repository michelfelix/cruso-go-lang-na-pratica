package main

import "fmt"

var a int = 1 // Escopo global sempre deve ser iniciado com palavras reservadas, como var, const, func, type, struct, interface, etc.

var texto1 , texto2 string = "Texto1", "Texto2" // Esse tipo de declaração só funciona com cariáveis do mesmo tipo

var texto3 = "Texto3"

func main() {

	//Início declaração multiplas variáveis
	var (
		b int
		c string
	)

	b, c, d := 10, "Vinte", true //Essa forma de declaração de variável só pode ser usada com uma variável nova e dentro de funções, fora delas não é permitido.

	texto2 = "Texto2 alterado" // Alterando o valor da variável texto2, que foi declarada no escopo global.

	texto3 := "Esse texto3 é diferente da variável texto3 declarada no escopo global" // Essa forma de declaração de variável só pode ser usada com uma variável nova e dentro de funções, fora delas não é permitido.

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(texto1)
	fmt.Println(texto2)
	fmt.Println(texto3) // Aqui será impresso o valor da variável texto3 declarada dentro da função main, e não a variável texto3 declarada no escopo global.

	//Fim declaração multiplas variáveis


}
