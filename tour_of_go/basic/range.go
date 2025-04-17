package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
		// range do laço for itera sobre uma slice ou map
		for i, v := range pow { // dois valores são retornados para cada iteração
		// o primeiro é o indice, o segundo uma copia do elemento daquele indice
				fmt.Printf("2**%d == %d\n", i, v)
		}

}
