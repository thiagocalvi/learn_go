package main

import "fmt"

func main() {
		var i, j int = 1, 2
		// Dentro de funções podemos utilizar declaração curta de variaveis
		k := 3
		// podo ser utilizado no lugar de var com tipo implicito
		c, python, java := true, false, "no!"

		fmt.Println(i, j, k, c, python, java)
}
