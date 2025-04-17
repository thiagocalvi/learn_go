package main

import "fmt"

const (
		// Uma constante sem tipo pega o seu tipo pelo contexto
		Big = 1 << 100
		Small = Big >> 99
)

func needInt(x int) int { x*10 + 1}
func needFloat(x float64) float64 { x * 0.1 }

func main() {
		fmt.Println(needInt(Small))
		fmt.Println(needFloat(Small))
		fmt.Println(needFloat(Big))
}
