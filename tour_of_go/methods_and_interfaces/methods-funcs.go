package main

import (
		"fmt"
		"math"
)

type Vertex struct {
		X, Y float64
}
		
func Abs(v Vertex) float64 {
		// metodos é apenas uma função com argumento receptor
		return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func main() {
		v := Vertex{3, 4}
		fmt.Println(Abs(v))
}
