package main

import "fmt"

func main() {
	var name string
	name = "Jonas"
	fmt.Print(name)
}

// Exemplo 1
//var name = Jonas (O proprio sistema ja entende que é uma string)

// Exemplo 2
// name := Jonas - A Nivel de package  nao podemos declarar desta forma se não da erro ao compilar
// Utilizar a forma do Exemplo 1, ou declarar a variável fora da func main
