//panic: erro do programador/erro execução tempo
//recover: ela interrompe o panic

package main

import "fmt"

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover: ", r)
		}
	}()
	panic("Panic!")
}
