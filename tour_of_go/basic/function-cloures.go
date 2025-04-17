package main

import "fmt"


// um closure é uma função valor que referencia variaveis de fora de seu corpo
func adder() func(int) int {
		sum := 0
		return func(x int) int {
				sum += x
				return sum
		}
}

func main() {
		pos, neg := adder(), adder()
		for i := 0; i < 10; i++ {
				fmt.Println(
						pos(i),
						neg(-2*i),
				)
		}
}
