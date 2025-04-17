package main

import "fmt"

// Striger é um tipo que pode descrever-se como uma string, muitos pacotes
// olham para essa interface para mostrar valores.

type Person struct {
		Name string
		Age int
}

func (p Person) String() string {
		return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
		a := Person{"Arthur Dent", 42}
		z := Person{"Zaphod Beeblebrox", 9001}
		fmt.Println(a, z)
}
