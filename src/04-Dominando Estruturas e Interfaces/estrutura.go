package main

import "fmt"

type pessoa struct {
	nome      string
	idade     int
	profissao string
}

func main() {
	fmt.Println(pessoa{"André", 36, "Programador"})
}

//são coleções de "campos"
//agrupar dados
//formar registros
