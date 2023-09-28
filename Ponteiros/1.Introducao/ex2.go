package main

import "fmt"

func main()  {
	var testValue string = "Otavio"
	fmt.Println(testValue)
	

	copyStringVALUE(testValue)
	fmt.Println("testValue sem mudar", testValue)

	originalStringVALUE(&testValue)
	fmt.Println("testValue mudado ", testValue)

}

func copyStringVALUE(stringValue string)  {
	stringValue = "Ingrid" //stringValue aqui é uma copia
	fmt.Println("copyStringVALUE ", stringValue)
}

func originalStringVALUE(stringValue *string)  {
	*stringValue = "Luana" //string value aqui é um ponteiro
  fmt.Println("originalStringVALUE ", *stringValue)
}






