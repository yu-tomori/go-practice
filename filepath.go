package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	// パスを結合する
	path := filepath.Join("dir", "main.go")
	// 拡張子を取る
	fmt.Println(filepath.Ext(path))
	// ファイル名を取得
	fmt.Println(filepath.Base(path))
	// ディレクトリ名を取得
	fmt.Println(filepath.Dir(path))
}
