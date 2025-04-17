package main

import "fmt"

func main() {
		// Alterando os dados de um slice modifica os elementos correspondentes da sua mateiz subjacente
		// Outros slices que partilham a mesma matriz vão ver essa mudança
		names := [4]string{
				"John",
				"Paul",
				"George",
				"Ringo",
		}

		fmt.Println(names)

		a := names[0:2]
		b := names[1:3]
		fmt.Println(a, b)

		b[0] = "XXX"
		fmt.Println(a, b)
		fmt.Println(names)
}
