package main

import (
	"fmt"
	"math"
)
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
		return fmt.SPrintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNEgativeSqrt(x)
	}
	return math.Sqrt(x), nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
