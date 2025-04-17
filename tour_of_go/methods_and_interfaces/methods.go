package main

import (
		"fmt"
		"math"
)

// go não tem classes, mas podemos definir metodos em tipos

type Vertex struct {
		X, Y float64
}
		
func (v Vertex) Abs() float64 {
     //^ o metodo é uma função com um argumento receptor especial
		return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func main() {
		v := Vertex{3, 4}
		fmt.Println(v.Abs())
}
