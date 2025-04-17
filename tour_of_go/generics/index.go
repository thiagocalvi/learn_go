package main

import "fmt"

// Funções podem ser escritas para funcionarem em varios tipos usando parametros de tipo, os paramentros de tipo de uma função aparecem entre colchetes, antes dos argumentos da função
func Index[T comparable](s []T, x T) int {
		// significa que s é um slice de qualquer tipo T que cumpre a restrição interna comparable, x é do mesmo tipo
		for i, v := range s {
				if v == x {
						return i
				}
		}

		return -1
		// comparable é uma restição que torna possivel usar os operadores == e !=
		//A função Index funciona para qualquer tipo que suporte comparação
}

func main() {
		si : []int{10, 20, 15, -10}
		fmt.Println(Index(si, 15))

		ss := []string{"foo", "bar", "baz"}
		fmt.Println(Index(ss, "hello"))
}
