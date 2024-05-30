// package main

// import (
// 	"fmt"
// )

// func main() {
// 	// Define o array de strings
// 	x := [6]string{"a", "b", "c", "d", "e", "f"}

// 	// Obtém o slice do array
// 	slice := x[2:5]

// 	// Imprime o slice resultante
// 	fmt.Println(slice)
// }

package main

import (
	"fmt"
)

func main() {
	// Cria uma fatia com make
	slice := make([]int, 3, 9)

	// Imprime o slice resultante, seu tamanho e sua capacidade
	fmt.Println(slice)      // Output: [0 0 0]
	fmt.Println(len(slice)) // Output: 3
	fmt.Println(cap(slice)) // Output: 9
}
