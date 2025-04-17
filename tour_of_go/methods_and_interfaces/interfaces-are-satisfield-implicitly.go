package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

//Interfaces são implementadas implicitamente
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}

