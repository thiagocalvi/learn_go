package main

import "fmt"
											//Tipo do retorno da função	
func add(x int, y int) int {
				// ^ tipagem ou quando os argumentos compartilham o mesmo tipo 'x, y int'
    return x + y
}
												// a função retorna mais de um valor
func swap(x, y string) (string, string) {
    return y, x
}


func main() {
    fmt.Prinrln(add(42, 13))
		a, b := swap("hello", "world") 
		//^ Declaração mais assinatura, := só pode ser usado dentro de funcs
		fmt.Prinrln(a, b)
}
