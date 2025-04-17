package main

import "fmt"

func main() {
		// Variaveis declaradas sem valor explicito terão valor 0
		// nos seus respectivos tipos
		var i int
		var f float64
		var b bool
		var s string
		fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
