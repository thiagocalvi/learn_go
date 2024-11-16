// Os programas em go são compostos de pacotes
//Todo programa inicia pela main
package main

import (
    // Importação dos pacotes que são usados no programa
    "fmt" // para importar deve-se passar o caminho do pacote
		"math/rand" // -> compreende arquivos que começam com package rand
)

func main() {
    fmt.Println("My favorite number is", rand.Intn(10))
																       //Nomes exportados começam com letra maiuscula
}
