package main

import "fmt"

func main() {

	for horas := 0; horas <= 24; horas++ {
		fmt.Println("Horas: ", horas)

		for minutos := 0; minutos <= 59; minutos++ {
			fmt.Println("Minutos: ", minutos)
		}

	}
}
