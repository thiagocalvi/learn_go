package main

// O import de pacotes podem ser feitos de duas formas
// forma consignada *recomendada
import (
	"fmt"
	"math"
)

// ou de forma individual
//import "fmt"
//import "math"

func main() {
	fmt.Printf("Now you have %g problems. \n", math.Sqrt(7))
}
