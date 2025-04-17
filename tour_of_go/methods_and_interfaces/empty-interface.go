package main

import "fmt"

func main() {
		var i interface{} // uma interface vazia pode conter valores de qualquer tipo desconhecido
		describe(i)

		i = 42
		describe(i)

		i = "hello"
		describe(i)
}

func describe(i interface{}) {
		fmt.Printf("(%v, %T)\n", i, i)
}
