package main

import (
		"fmt"
		"math"
)

type Vertex struct {
		X, Y float64
}

// receptor com valor em copia
func (v Vertex) Abs() float64 {
		return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

// receptor sendo ponteiro
// o tipo do receptor tem a sintaxe *t para um tipo t, t deve ser um tipo local
func (v *Vertex) Scale(f float64) {
		v.X = v.X * f
		v.Y = v.Y * f
}

func main() {
		v := Vertex{3, 4}
		v.Scale(10)
		fmt.Println(v.Abs())
}
