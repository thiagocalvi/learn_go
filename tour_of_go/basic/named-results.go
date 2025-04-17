package main

import "fmt"

// Valores de retorno podem ser nomeados e agirem como variaveis
func split(sum int) (x, y int) {
		// Os nomes são usados para documentar
		// o significado dos valores de retorno
		x = sum * 4 / 9
		y = sum - x
		return
		// return sem argumento retorna os valores atuais dos resultados
		// Recomendado usar em funções curtas
}

func main() {
		fmt.Println(split(17))
}
