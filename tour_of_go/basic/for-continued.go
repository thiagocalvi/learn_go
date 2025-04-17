package main

import "fmt"

func main() {
		sum := 1
		// instrução inicial e pós-instrução são opcionais
		for ; sum < 1000; {
				sum += sum
		}

		fmt.Println(sum)
}
