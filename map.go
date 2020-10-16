package main

func main() {
    // ゼロ値はnil
    // var m map[string]int
    // makeで初期化
    // m = make(map[string]int)
    // 容量を指定できる
    // m = make(map[string]int, 10)
    // 空の場合
    // m2 := map[string]int{}

    // リテラルで初期化
    m1 := map[string]int{"x": 10, "y": 20}
    // キーを指定してアクセス
    println(m1["x"])
    // キーを指定して入力
    m1["z"] = 30
    // 存在を確認する
    n, ok := m1["z"]
    println(n, ok)
    // キーを指定して削除する
    delete(m1, "z")
    // 削除されていることを確認
    n, ok = m1["z"]
    println(n, ok)
}
