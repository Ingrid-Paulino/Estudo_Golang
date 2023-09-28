package main

import "fmt"

func assertValue(value1, value2 string) {
	if value1 == value2 {
		fmt.Print("São iguais")
	}

	panic("Não são iguais")
}

func main()  {
	defer func ()  { //funcao anonima com defer, no fim da execucao do codigo vai comferir se teve algum panic
		//recover é como o try/cath de javascript
		if err := recover(); err != nil {
			fmt.Println("Recebi erro de um panic, recuperando execução. Erro:", err)
			//Aqui poderiamos fechar conecxoes se tivesse alguma em aberto
			//fecha listener com filas
		}
	}()
	assertValue("hun", "coding")
}