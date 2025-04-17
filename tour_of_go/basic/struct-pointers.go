package main

import "fmt"

type Vertex struct {
		X int
		Y int
}

func main() {
		v := Vertex{1, 2}
		// campos de estruturas podem ser acessados atraves de ponteiros de estrutura
		p := &v
		p.X = 1e9
		fmt.Println(v)
}
