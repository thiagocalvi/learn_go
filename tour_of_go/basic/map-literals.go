package main

import "fmt"

type Vertex struct{
		Lat, Long float64
}


func main() {
		// maps literais são como estruturas literais, mas as chaves são obrigatorias
		m := map[string]Vertex{
				"Bell Labs": Vertex{
				40.68433, -74.39967,
		},
				"Google": Vertex{
				37.42202, -122.08408,
				}, 
		
		}

		fmt.Println(m)
}
