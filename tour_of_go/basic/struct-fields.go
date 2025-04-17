package main

import "fmt"

type Vertex struct {
		X int
		Y int
}

func main() {
		v := Vertex{1, 2}
		v.X = 4 // os campos de uma estrtutura são acessados com .
		fmt.Println(v.X)
}
