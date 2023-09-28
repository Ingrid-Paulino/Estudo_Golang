package main

import "fmt"

// Com o type assertion vc consegue saber qual é o objeto que implementa determida interface
// %T : informa qual objeto esta implementando a interface
//%v : printa os valores do objeto que a interface recebeu
// %#v : printa chave valor do objeto que a interface recebeu

type huncoding interface {
	printName()
}

type otavio struct {
	name string
	test string
}

type test struct{}


func (o otavio) printName()  { //A interface huncoding esta sendo implementada pelo objeto otavio
	fmt.Println("Hello World")
}


func Print(hun huncoding)  {
	//Printa quem implementa a interface huncoding
	fmt.Printf("O objeto por tras desta interface huncoding é o objeto: %T \n", hun )

	//transforma a interface em objeto para que eu cosiga pegar os valores do objeto dentro dela
	obj := hun.(otavio) //hun é a interface e eu passo nos parentes o nome do objeto que eu quero transformar
	 //Consigo ter acesso aos valores
	  fmt.Printf("%s, %s \n", obj.name, obj.test)

	hun.printName()
}

func main() {
	obj := otavio{"ingrid", "Passou parabéns"}
	Print(obj)
}