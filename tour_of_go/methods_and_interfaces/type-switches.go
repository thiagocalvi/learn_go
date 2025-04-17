package main

import "fmt"

func do(i interface{}) {
		// type switch permite diversas type assertions
		switch v := i.(type) {
		// os cases especificam os tipos, esse valores são comparados com o tipo dos valores 
		// determinados da interface informada
		case int:
				fmt.Printf("Twice %v is %v\n", v, v*2)
		case string: 	
				fmt.Printf("%q is %v bytes long\n", v, len(v))

		default:
				fmt.Printf("I don't know about type %T\n", v)
		}
}

func main() {
		do(21)
		do("hello")
		do(true)
}
