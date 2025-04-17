package main

import "fmt"

func main() {
		// [n]T é uma matriz de n valor do tipo T
		// matrizes não podem ser redimensionadas
		var a [2]string
		a[0] = "Hello"
		a[1] = "World"
		fmt.Println(a[0], a[1])
		fmt.Println(a)

		// cria uma matriz e a inicializa com valores
		primes := [6]int{2, 3, 5, 7, 11, 13}
		fmt.Println(primes)
}
