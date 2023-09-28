package main

//Type switch valida se o obeto da interface corresponde com o que esta sendo passado
// caso não tenha validado e os objetos forem diferentes, a aplicaçao dara penic
// O interesante é tratar esse panic 

import "fmt"

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

func (t test) printName() {
	fmt.Println("Hello World")
}




func Print(hun huncoding)  {
	//Se for passado um objeto para conversão que nao espero, dara penic no projeto. Pra isso nao ocorrer podemos tratar/validar o erro 
	otavioObj, ok := hun.(otavio) //Ao inves de otavio deveria ser test
	if !ok { //evita dar panic a tratativa de erro
		//fmt.Print("Error trying to convert")
		return 
	}

	otavioObj.printName()
}

func main() {
	//obj := otavio{"ingrid", "Passou parabéns"}
	obj := test{}

	Print(obj)
}