package main

import "fmt"

type retangulo struct {
	comprimento, altura int
}

func (r retangulo) area() int {
	return r.comprimento * r.altura
}

func (r retangulo) perimetro() int {
	return 2*r.comprimento + 2*r.altura
}

// Método com receptor ponteiro para modificar o objeto
func (r *retangulo) dobrarAltura() {
	r.altura *= 2
}

func main() {
	r := retangulo{comprimento: 10, altura: 5}

	fmt.Println("Área inicial:", r.area())           // Saída: Área inicial: 50
	fmt.Println("Perímetro inicial:", r.perimetro()) // Saída: Perímetro inicial: 30

	r.dobrarAltura() // Modifica a altura original

	fmt.Println("Nova área:", r.area())           // Saída: Nova área: 100
	fmt.Println("Novo perímetro:", r.perimetro()) // Saída: Novo perímetro: 40
	fmt.Println("Nova altura:", r.altura)         // Saída: Nova altura: 10
}
