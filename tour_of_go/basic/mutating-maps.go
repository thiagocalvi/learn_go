package main

import "fmt"

func main() {
		m := make(map[string]int)

		// inserir ou atualizar um elemento no map
		m["Answer"] = 42
		fmt.Println("The value:", m["Answer"]) // m["Answer"] recupera um elemento

		m["Answer"] = 48
		fmt.Println("The value", m["Answer"])

		delete(m, "Answer") // excluir um elemento
		fmt.Println("The value:", m["Answer"])

		v, ok := m["Answer"] // teste que uma chave esta presente
		// v é o valor na chave se estiver presente, ok é um bool
		fmt.Println("The value:", v, "Present?", ok)
}
