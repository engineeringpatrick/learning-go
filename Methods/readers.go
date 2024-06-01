package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (reader MyReader) Read(b []byte) (int, error) {
	// for i := range b {
	// 	b[i] = 'A'
	// }
	// return len(b), nil
	b[0] = 'A'
	return 1, nil
}

func main() {
	reader.Validate(MyReader{})
}
