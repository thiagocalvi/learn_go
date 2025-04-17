package main

import (
		"fmt"
		"math"
)

func main() {
		var x, y int = 3, 4
		var f float64 = math.Sqrt(float64(x*x + y*y))
		var z uint = uint(f) // a espressão T(v) converte o valor v para o tipo T
		// atribuição entre itens de tipos diferente requer conversão explicita 
		fmt.Println(x, y, z)
}
