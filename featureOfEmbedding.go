package main

import "fmt"

// 型リテラルでなければ埋め込められる
// ・typeで定義したものや組み込み型
// ・インタフェースも埋め込められる
// インタフェースの実装
// ・埋め込んだ値のメソッドもカウント
type Stringer interface {
	String() string
}

type Hex int

func (h Hex) String() string {
	return fmt.Sprintln("%x", int(h))
}

type Hex2 struct{ Hex }
