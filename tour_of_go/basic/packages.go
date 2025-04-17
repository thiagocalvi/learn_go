// Programas go são compostos de pacotes
package main // todo programa se inicia pelo pacote main

import (
	// esses são pacotes importados
	"fmt"
	"math/rand" // caminho de importação do pacote
	// o nome do pacote é mesmo que o ultimo elemento do caminho de importação
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}
