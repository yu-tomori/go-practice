package main

import "fmt"

// インタフェースをインタフェースに埋め込む
// ・複数のインタフェースを合成する
// ・複雑なインタフェースが必要な場合
type Reader interface {
	Read(p []byte) (n int, err error)
}
type (
	Writer interface {
		Write(p []byte) (n int, err error)
	}
)

type ReadWriter interface {
	Reader
	Writer
}
