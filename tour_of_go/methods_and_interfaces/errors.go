package main

import (
		"fmt"
		"time"
)

type MyError struct {
		When time.Time
		What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)	
}

// Programas go expressão estados de erro com valores error, o tipo error é uma inteface embutida
func run() error {
		return &MyError{
				time.Now(),
				"it didn't work",
		}
}

func main() {
		// Funções frequentemente retornam um valor error e a chamada do código deve lidar com erros testando se o erro é igual a nil  
		if err := run(): err != nil {
				fmt.Println(err)
		}
}
