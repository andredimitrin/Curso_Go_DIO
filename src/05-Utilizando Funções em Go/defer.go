//defer: escalona nossas funções após o final do programa

package main

import "fmt"

func dia1() {
	fmt.Println("Domingo")
}

func dia2() {
	fmt.Println("Segunda")
}

func main() {
	defer dia2()
	dia1()
}
