//função também é chamada de procedimento ou sub-rotina
//parte do código
//ela pega um dado de entrada e dá uma dado de saída

package main

import "fmt"

func media(lista []float64) float64 { //recebe uma lista de float64 e retorna um float64
	total := 0.0
	for _, valor := range lista {
		total += valor
	}
	return total / float64(len(lista))
}

func main() { // programa que calcula a media de uma sala
	lista := []float64{98, 93, 77, 50, 80} //lista de notas

	fmt.Println(media(lista))

	// total := 0.0
	// for _, valor := range lista {
	// 	total += valor
	// }

	// fmt.Println(total / float64(len(lista))) //imprime a media

}
