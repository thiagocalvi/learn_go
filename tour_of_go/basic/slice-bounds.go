package main

import "fmt"

func main() {
		s := []int{2, 3, 5, 7, 11, 13}

		// ao fatiar podemos omitir os limites altas e baixas
		s = s[1:4]
		fmt.Println(s)

		// o padrão é igual a zero para o limite baixo e o comprimento
		// da slice para o limite alto
		s = s[:2]
		fmt.Println(s)

		s = s[1:]
		fmt.Println(s)
}
