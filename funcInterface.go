package main

import "fmt"

type Func func() string

func (f Func) String() string { return f() }

func main() {
	var s fmt.Stringer = Func(func() string { // Func型へのキャスト(型変換)
		return "hi" // 無名関数=関数リテラル
	})
	fmt.Println(s)
}
