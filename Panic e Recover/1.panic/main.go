package main

import "fmt"

func assertValue(value1, value2 string) {
	if value1 == value2 {
		fmt.Print("São iguais")
	}

	panic("Não são iguais")
}

func main()  {
	assertValue("hun", "coding")
}