package main

import (
	"errors"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (reader *rot13Reader) Read(b []byte) (int, error) {
	_, err := reader.r.Read(b)

	if err != nil {
		return 0, errors.New("Something happened when reading original byte stream")
	}

	for i, _ := range b {
		if b[i] >= 'a' && b[i] <= 'z' {
			b[i] = 'a' + ((b[i] - 'a' + 13) % 26)
		}
		if b[i] >= 'A' && b[i] <= 'Z' {
			b[i] = 'A' + ((b[i] - 'A' + 13) % 26)
		}
	}

	return len(b), nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
