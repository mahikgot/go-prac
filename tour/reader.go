package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (myReader MyReader) Read(b []byte) (int, error) {
	b[0] = byte(65)
	return 1, nil
}

func main() {
	reader.Validate(MyReader{})
}
