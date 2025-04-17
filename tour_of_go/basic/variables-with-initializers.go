package main

import "fmt"

// var pode incluir incializadores, um por variavel
var i, j int = 1, 2
// se o valor inicial estiver presente o tipo pode ser omitido

func main() {
		var c, python, java = true, false, "no!"
		fmt.Println(i, j, c, python, java)
}
