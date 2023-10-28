package webpfuzz

import (
	"bytes"
	"golang.org/x/image/webp"
)

func Fuzz(data []byte) int {
	r := bytes.NewReader(data)
	_, err := webp.Decode(r)
	if err != nil {
		return 0
	}
	return 1
}
