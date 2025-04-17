package main 

import "fmt"

// um tipo pode ser parametrizado com um parametro de tipo
type List[T any] struct {
		next *List[T]
		val T
}

func main() {
		list := List[int]{}
		fmt.Println(list)
}
