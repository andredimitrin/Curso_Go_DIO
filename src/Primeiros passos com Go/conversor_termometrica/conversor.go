package main

import (
	"fmt"
)

func kelvinToCelsius(kelvin float64) float64 {
	return kelvin - 273.15
}

func celsiusToKelvin(celsius float64) float64 {
	return celsius + 273.15
}

func main() {
	// Temperatura em Kelvin
	temperaturaKelvin := 373.15

	// Convertendo de Kelvin para Celsius
	temperaturaCelsius := kelvinToCelsius(temperaturaKelvin)

	fmt.Printf("%.2f Kelvin é equivalente a %.2f Celsius\n", temperaturaKelvin, temperaturaCelsius)

	// Temperatura em Celsius
	temperaturaCelsius2 := 100.0

	// Convertendo de Celsius para Kelvin
	temperaturaKelvin2 := celsiusToKelvin(temperaturaCelsius2)

	fmt.Printf("%.2f Celsius é equivalente a %.2f Kelvin\n", temperaturaCelsius2, temperaturaKelvin2)
}
