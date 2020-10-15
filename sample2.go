// iota
// 連続した定数を作るための仕組み
// グループ化された名前付き定数の定義で使われる
// 0から始まり1ずつ加算される値として扱われる
package main
import (
    "fmt"
)

func main() {
    const (
        a = iota
        b
    )
    const (
        c = 1 << iota
        d
        e
    )
    fmt.Println(a, b, c, d, e)
}
