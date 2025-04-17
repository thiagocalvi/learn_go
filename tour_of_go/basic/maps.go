package main

import "fmt"

type Vertex struct{
		Lat, Long float64
}

// Um map mapeia chaves para valores
var m map[string]Vertex // o valor zero de um map é nil, map nil tem chaves, nem chaves podem ser adicionadas

func main() {
		m = make(map[string]Vertex) // retorna um map inicializado, com determindao tipo
		m["Bell Labs"] = Vertex{40.68433, -74.39967}
		fmt.Println(m["Bell Labs"])
}
