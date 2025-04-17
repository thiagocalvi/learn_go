package main

import "fmt"

type I interface {
		M()
}

// um valor de interface nil detém nem o valor nom tipo concreto
func main() {
		var i I
		describe(i)
		i.M() // chamar o metodo resulta em um erro
}

func describe(i I) {
		fmt.Printf("(%v, %T)\n", i, i)
}
