package main

import "fmt"

// Ponteiro permite que compartilhemos valores em threads diferentes
// Se vc passa uma valor pra uma função sem ponteiro, vc estara passando uma copia
// Ponteiro em Golang é uma variável que é usada para armazenar o endereço de memória de outra variável.
//Podemos passar uma copia da variavel quando não queremos alterar o valor real

//Operadores: * e &
// Em go um ponteiro é representado pelo caracter *(asterisco)
// * tbm é usado para desreferenciar variaveis de ponteiro. Cancelar a referencia de um ponteiro nós dá acesso ao valor para o qual o ponteiro aponta.
//ex:
// var x int = 10
// var y *int = &x  - y aponta pra o endereço de memoria x
//*y = 12 - o x vai passar a valer 12
//fmt.println(*y)
//Usamos o operador & para encontrar o endereço de uma variável.

//Quando instaciamos alguma variavel com ponteiro, interface, slices, channels maps e sem atribuir um valor, o valor inicial é null/nil(nulo)
//ex:
//var pointer *string
// fmt.println(pointer) //nil

		 func main()  {
			var x int = 100
			var y *int = &x
			fmt.Println(x, y)  //y esta ira retornar o endereço de memoria do x
			fmt.Println(&x)  //retorna o endereço de memoria de x
			fmt.Println(*y)  //com ponteiro consigo pegar o valor dentro de x
			fmt.Println(&y)  //com & vai retornar o endereço de memoria que y esta guardado
		 }
