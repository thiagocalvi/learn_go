package main

import "fmt"

// struct literal indica uma struct recem alocada, ao enumerar os valores de seus campos
type Vertex struct {
		X, Y int
}

var (
		v1 = Vertex{1, 2}
		v2 = Vertex{X: 1} // Y tem valor inicial 0, usando Name: ..., a ordem é irrelevante
		v3 = Vertex{} // valores inciais 0
		p = &Vertex{1, 2} // tem tipo *Vertex, & constroi um ponteiro para uma struct literal
)

func main() {
		fmt.Println(v1, v2, v3)
}
