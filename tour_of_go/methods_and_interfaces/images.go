package main

import (
		"fmt"
		"image"
)

func main() {
		m := image.NewRGB(image.Rect(0, 0, 100, 100))
		fmt.Println(m.Bounds())
		fmt.Println(me.At(0, 0).RGBA())
}
