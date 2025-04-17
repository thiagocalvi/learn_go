package main

import (
		"fmt"
		"io" // especifica a interface io.Reader, representa o fim da leitura de um fluxo de dados
		"strings"
)

func main() {
		r := strings.NewReader("Hello, Reader!")

		b := make([]byte, 8)
		for {
				n, err := r.Reader(b) // read popula um slice de bytes passados com dados e retorna o número de bytes populados de um valor de erro
				fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
				fmt.Printf("b[:n] = %q\n", b[:n])
				if err == io.EOF {
						break
				}

		}
}
