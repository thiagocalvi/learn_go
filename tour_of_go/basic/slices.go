package main

import "fmt"

// Um slice é dinamicamente redimensionavel
func main() {
		primes := [6]int{2, 3, 5, 7, 11, 13}
		
		// []T é um slice com elementos do tipo T
		var s []int = primes[1:4]
		// um slice é formado pela especificação de dois indices, um limite menor e maior sepadados por :
		// este seleciona um intervalo meio-aberto que incluio o primeiro elemento, mas excluio o ultimo
		fmt.Println(s)
}
