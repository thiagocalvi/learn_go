package main

import "fmt"

func main() {
		var i interface{} = "hello"

		s := i.(string) // type assertion fornece acesso ao valor concreto subjacente de um valor de interface
		fmt.Println(s)

		s, ok := i.(string)
		fmt.Println(s, ok)

		f, ok := i.(float64)
		fmt.Println(f, ok)

		f = i.(float64) // panic
		fmt.Println(f)

}
