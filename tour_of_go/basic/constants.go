package main

import "fmt"

const Pi = 3.14

func main() {
		const World = "@" // constantes não pode ser declaradas usando :=
		fmt.Println("Hello", World)
		fmt.Println("Happy", Pi, "Day")

		const Truth = true
		fmt.Println("Go rules?", Truth)
}
