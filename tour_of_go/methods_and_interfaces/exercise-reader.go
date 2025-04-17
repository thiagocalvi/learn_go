package main

import "golang.org/x/tour/reader"

type MyReader struct {}

func (r MyReader) Reader(p []byte) (int, error) {
		for i := range p {
				p[i] = "A" 
		}

		return len(p), nil 
}

func main() {
		reader.Validate(MyReader{})
}
