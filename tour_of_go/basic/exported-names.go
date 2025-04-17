package main

// nomes exportados são aqueles que começar com letra maiuscula
// quando importamos um pacote, só podemos referenciar seus nomes exportados
// os nomes não exportodos não são acessiveis fora do pacote
import (
	"fmt"
	"math"
)

func main() {
		// Println e Pi são nomes exportados
		fmt.Println(math.Pi)
}
