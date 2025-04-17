package main

import (
		"fmt"
		"time"
)

func say(s string) {
		for i := 0; i < 5; i++ {
				time.Sleep(100 * time.Millisecond)
				fmt.Println(s)
		}
}

func main() {
		// uma goroutine é um segmento leve e gerenciado pelo runtime de go
		go say("world") //inicia uma nova execução goroutine 
		say("hello")
}

// goroutines executam no mesmo espaço de endereço, para que o acesso a memoria compartilhada seja sincronizado
