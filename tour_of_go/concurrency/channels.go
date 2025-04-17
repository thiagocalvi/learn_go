package main

import "fmt"

func sum(s []int, c chan int) {
		sum := 0
		for _, v := range s {
				sum += v
		}

		c <- sum
}

// canais são um conduto tipado que atrves dele podemos enviar e receber valores com o operador de canal <-, os dados fluem na direção da seta

func main() {
		s := []int{7, 2, 8, -9, 4, 0}

		c := make(chan int) // os canais devem ser criados antes de usar
		go sum (s[:len(s)/2], c)
		go sum (s[len(s)/2:], c)

		x, y := <- c, <- c

		fmt.Println(x, y)
}
