package main

import (
	"fmt"
	"math"
)

// Geometria: Interface para formas geométricas
type Geometria interface {
	Area() float64
	Perimetro() float64
}

// Quadrado: Estrutura para representar um quadrado
type Quadrado struct {
	Lado float64
}

// Circulo: Estrutura para representar um círculo
type Circulo struct {
	Raio float64
}

// Area: Método para calcular a área do quadrado (Lado * Lado)
func (q Quadrado) Area() float64 {
	return q.Lado * q.Lado
}

// Perimetro: Método para calcular o perímetro do quadrado (4 * Lado)
func (q Quadrado) Perimetro() float64 {
	return 4 * q.Lado
}

// Area: Método para calcular a área do círculo (π * Raio²)
func (c Circulo) Area() float64 {
	return math.Pi * c.Raio * c.Raio
}

// Perimetro: Método para calcular o perímetro (circunferência) do círculo (2 * π * Raio)
func (c Circulo) Perimetro() float64 {
	return 2 * math.Pi * c.Raio
}

// Medir: Função genérica para exibir informações de uma forma geométrica
func Medir(g Geometria) {
	fmt.Println("Forma:", g)                       // Imprime a representação da forma (string)
	fmt.Printf("Área: %.2f\n", g.Area())           // Formatação para duas casas decimais
	fmt.Printf("Perímetro: %.2f\n", g.Perimetro()) // Formatação para duas casas decimais
}

func main() {
	q := Quadrado{Lado: 10}
	c := Circulo{Raio: 5}

	fmt.Println("---- Quadrado ----")
	Medir(q)
	fmt.Println("---- Círculo ----")
	Medir(c)
}
