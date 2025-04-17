package main

import "fmt"

func main() {
		var s []int // O valor 0 de um slice é nil
		fmt.Println(s, len(s), cap(s))
		if s == nil {
				fmt.Println("nil!")
		}
}
