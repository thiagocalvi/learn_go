package main

import "fmt"

// uma função pode retornar qualquer número de resultados
func swap(x, y string) (string, string) {
		return y, x
}

func main() {
		a, b := swap("hello", "world")
		fmt.Println(a, b)
}
