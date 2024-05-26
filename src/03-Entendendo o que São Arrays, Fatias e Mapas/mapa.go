package main

import "fmt"

func main() {

	// x := make(map[string]int)
	// x["key"] = 10
	// fmt.Println(x["key"])

	// x := make(map[int]int)
	// x[1] = 20
	// x[2] = 30
	// fmt.Println(x[1], x[2])

	elemento := make(map[string]string)
	elemento["nome"] = "André"
	elemento["sobrenome"] = "Dimitrin"

	fmt.Println(elemento["nome"], elemento["sobrenome"])
}

//coleção ordenada de pares chave-valor
//var x map [string]int
// x é um mapa de strings para ints
