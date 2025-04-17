package main

import "fmt"

func main() {
		// make criar matrizes dinamincamente
		a := make([]int, 5) // aloca uma mateiz zerada e retorna um slice que se refere a essa matriz. 5 é o len. len = 5
		printSlice("a", a)

		b := make([]int, 0, 5) // len = 0, cap = 5
		printSlice("b", b)

		c := b[:2]
		printSlice("c", c)

		d := c[2:5]
		printSlice("d", d)
}

func printSlice(s string, x []int) {
		fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
