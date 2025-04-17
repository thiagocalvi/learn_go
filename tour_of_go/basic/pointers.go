package main

import "fmt"

func main() {
		// Um ponteiro guarda na memoria o endereço de uma variavel
		// definir um ponteiro para um tipo-> var p *int

		i, j := 42, 2701

		p := &i //ponteiro para i, & gera um ponteiro para seu operando
		fmt.Println(*p) //lê i atraves do ponteiro p
		*p = 21 // define i atraves do ponteiro p

		p = &j
		*p = *p / 37
		fmt.Println(j)

		// Go não faz aritimética de ponteiros

}
