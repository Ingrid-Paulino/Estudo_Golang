package main

//Para performace é melhor deixar definido o tipo da chave valor que vc espera no map
//Dessa forma o codigo executa mais rapido
//forma melhor
var test1 map[string]string = make(map[string]string)

//Nesse caso valor pode ser qualquer coisa
//interface é um tipo any em golang
//Devemos usar map de string com interface, quando não sabemos o que vem no json
var test2 map[string]interface{} = make(map[string]interface{})
