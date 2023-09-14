// Declarar meu pacote principal
package main

import "fmt"

// Declarar variáveis constantes
const ebulicaiFahrenheit float64 = 212.0

func main() {
	// Declarar e inicializar variáveis
	// tempFF := 32.0
	// tempFC := 0.0

	// // Calcular a temperatura em Celsius
	// tempF := ebulicaiFahrenheit
	// tempC := (tempF - 32) * 5 / 9

	// // Imprimir os resultados formatados
	// fmt.Printf("A temperatura em fahrenheit é: %g\n", tempF)
	// fmt.Printf("A temperatura em celsius é: %g\n", tempC)

	// // Imprimir os tipos das variáveis
	// fmt.Printf("O tipo de tempFF é: %T\n", tempFF)
	// fmt.Printf("O tipo de tempFC é: %T\n", tempFC)
	// fmt.Printf("O tipo de tempF é: %T\n", tempF)
	// fmt.Printf("O tipo de tempC é: %T\n", tempC)

	  tempF = 212.0 
      tempC = (tempF – 32)*5/9 
      fmt.Printf(“A temperatura em °F = %g e a temperatura em °C = %g.”, tempF, tempC) 
}
