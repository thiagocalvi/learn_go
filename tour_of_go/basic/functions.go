package main

import "fmt"

// Uma função pode ter 0 ou mais parametros
func add(x int, y int) int {
		// o nome parametro é seguido do seu tipo 
		return x + y
}

func add2(x, y int) {
		//quando dois ou mais parâmetros nomeados consecutivos de uma função compartilham o mesmo tipo, podes omitir o tipo de todos menos o último	
		return x + y
}

func main() {
		fmt.Println(add(4, 31))
		fmt.Println(add2(4, 13))
}
