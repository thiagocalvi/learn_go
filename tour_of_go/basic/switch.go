package main

import (
		"fmt"
		"runtime"
)

func main() {
		fmt.Print("GO runs on ")
		switch os := runtime.GOOS; os {
				case "darwin":
						fmt.Println("OS X.")
				case "linux":
						fmt.Println("Linux.")
				default:
						fmt.Println("%s.\n", os)
		}
		// break é fornecido automaticamente
}
