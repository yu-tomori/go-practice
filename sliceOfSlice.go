package main
import "fmt"

func main() {
    ns := []int{10, 20, 30, 40, 50}
    n, m := 2, 4

    // n番目以降のスライスを取得する
    fmt.Println(ns[n:])

    // 先頭からm-1番目までのスライスを取得する
    fmt.Println(ns[:m])

    // capを指定する
    ms := ns[:m:m]
    fmt.Println(cap(ms))

    // merge slices
    a := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
    b := []int{40, 50}
    c := append(a, b...)
    fmt.Println(c)

    // cut slice
    d := append(a[:2], a[5:]...)
    fmt.Println(d)

    // delete slice
    e := append(a[:2], a[2+1:]...)
    fmt.Println(e)

    e = e[:2+copy(e[2:], e[2+1:])]
    fmt.Println(e)
}
