package main

import "fmt"

func main() {
		// declarar uma variavel sem especificar o tipo
		// ela tera seu tipo inferido a partir do valor do lado direito
		v := 4.1
		// quando o lado direito contem uma constante numerico 
		// o tipo dependa da precisão da constante
		fmt.Printf("v is of type %T\n", v)
}
