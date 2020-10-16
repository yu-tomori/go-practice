package main
import "fmt"

func main() {
    // ゼロ値で初期化
    var ns1 [5]int
    // 配列リテラルで初期化
    var ns2 = [5]int{10, 20, 30, 40, 50}
    // 要素数を型から推論
    ns3 := [...]int{10, 20, 30, 40, 50}
    // 5番目が50, 10番目が100で他が0の要素数11の配列
    ns4 := [...]int{5: 50, 10: 100}
    println(ns1[0], ns2[0], ns3[0], ns4[0])

    ns := [...]int{10, 20, 30, 40, 50}
    // 要素にアクセス
    println(ns[3]) // 添字は変数でも良い
    // 長さ
    println(len(ns))
    // スライス演算
    fmt.Println(ns[1:3])
}
