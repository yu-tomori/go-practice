package main

import "fmt"

type Hoge struct {
	N int
}

type Fuga struct {
	// 構造体に匿名フィールドを埋め込む機能
	Hoge // 名前のないフィールドになる
}

func main() {
	f := Fuga{Hoge{N: 100}}

	// Hoge型のフィールドにアクセスできる
	fmt.Println(f.N)

	// 型名を指定してアクセスできる
	fmt.Println(f.Hoge.N)
}
