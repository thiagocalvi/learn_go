package main

import "fmt"

func main() {
		fmt.Println("counting")
		// chamadas de funções adiadas são empilhadas em uma pilha
		for i := 0; i < 10; i++ {
				defer fmt.Println(i)
		}
		// quando a função retorna, as chamdas de adiamento são desempilhadas e executadas 

		fmt.Println("done")
}
