package main

import "fmt"

func main() {
		pow := make([]int, 10)
		
		//Podemos ignorar valor de retorno
		for i := range pow {
				pow[i] = 1 << uint(i) // 2**i
		}

		for _, value := range pow {
				fmt.Printf("%d\n", value)
		}
}
