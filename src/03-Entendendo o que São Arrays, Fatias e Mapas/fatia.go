package main

import (
	"fmt"
)

func main() {

	fatia1 := []int{1, 2, 3}
	fatia2 := make([]int, 1)
	copy(fatia2, fatia1)
	fmt.Println(fatia1, fatia2)

}

// var x []float64
//slice:=make([]float64,4)
//slice := [low:high]
//slice := arr [0:5]
//acrescentar elemento APPEND
//copiar slice COPY
