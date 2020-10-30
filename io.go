package main

import (
	"fmt"
	"os"
)

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

func main() {
	fmt.Fprintf(os.Stderr, "エラー")
}
