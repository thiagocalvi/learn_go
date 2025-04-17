package main

import (
		"fmt"
		"math"
)

func Sqrt(x float64) float64 {
		z := 1.0
		for i := 1; i <= 100; i++ {
				z -= (z * z - x) / (2 * z)
				fmt.Println("Valor de z: ", z)
		}

		return z
}

func main() {
		value := 2.0
		fmt.Println("Valor calculado por Sqrt: ", Sqrt(value))
		fmt.Println("Valor calculado por math/Sqrt: ", math.Sqrt(value))
}
