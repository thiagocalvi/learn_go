package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
		if t == nil {
				fmt.Println("<nil>")
				return
		}	
		fmt.Println(t.S)
}

func main() {
	var i I

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"Hello"}
	describe(i)
	i.M()
}

func describe(i I) {
		fmt.Println("(%v, %T)\n", i, i)
}

