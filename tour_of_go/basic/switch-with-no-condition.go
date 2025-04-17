package main

import (
		"fmt"
		"time"
)

func main() {
		t := time.Now()
		// switch sem uma condiçaõ é o mesmo que switch true
		switch {
		case t.Hour() < 12:
				fmt.Println("Good morning!")
		case t.Hour() < 17:
				fmt.Println("Goof afternoon.")
		default:
				fmt.Println("Good evening.")
		}

}
