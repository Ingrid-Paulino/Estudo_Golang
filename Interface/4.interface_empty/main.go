package main

import "fmt"

//Nem sempre interface terao tipos e funcoes definidas.
//Nem sempre as interfaces são definidas para implementar methodos

//Tipo 1
//É necessário implementar dois mérodos para "satisfazer" esta interface
type shape interface {
	area() float64
	perimeter() float64
}

//Tipo 2
//É necessário implementar apenas um mérodos para "satisfazer" esta interface
type shape2 interface {
	area() float64
}

//Tipo 3 
//Qualquer objeto implementa esta interface, por ela não pedir nada
type shape3 interface {}



// exemplo:

type huncoding interface {} //Por a interface ser vazia, ela não pede nenhum objeto, entt posso usar qualquer objeto nessa interface

type test struct{
	name string
} //esse obj não implementa nenhuma função

func Print(hun huncoding)  { //posso passar qualquer objeto por parametro, pois huncoding não implementa/pede nada
	fmt.Printf("Hello %s \n", hun)
  fmt.Printf("Hello %s \n", hun.(test).name) //Se eu passar outro obj diferente de test dara panic o sistema, o interesante seria tratar o hun.(test).name como feito na pasta 3.type_switc
}

func main() {

  var testInterfaceEmpty interface{} //Essa variavel pode receber qualquer valor
	testInterfaceEmpty = "Hello"
	testInterfaceEmpty = 2
	testInterfaceEmpty = 2.5
	testInterfaceEmpty = true

	var testAny any // o tipo any em golang é a mesma coisa de interface{}
	testAny = true
	testAny = "Hello"
	testAny = 2
	testAny =2.5

	obj := test{name: "Ingrid"}
	Print(obj)
	fmt.Print(testInterfaceEmpty, "\n")
	fmt.Print(testAny)
}