package main

import (
	"fmt"
	"os"
)

func reader() {
	// 読み込み時にファイルを開く
	sf, err := os.Open(src)
	if err != nil {
		return err
	}
	// 関数終了時に閉じる
	defer sf.Close()
}

func writer() {
	// 書き込み時にファイルを開く
	df, err := os.Create(dst)
	if err != nil {
		return err
	}
	// 関数終了時に閉じる
	defer func() {
		if err := df.Close(); err != nil {
			rerr = err
		}
	}()
}

func main() {
	reader()
	writer()
}
