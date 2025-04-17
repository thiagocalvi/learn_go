package main

import "fmt"

func main() {
		//defer adia a execução de uma função até o final do retorno da função
		defer fmt.Println("world")
		// os argumentos das chamadas adiada são avaliados imediatamente, mas a função chamda não é executada até o retorno da função

		fmt.Println("hello")
}
