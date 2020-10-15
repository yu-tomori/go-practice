package main
import "fmt"

func main() {
    for i := 0; i <= 10; i++ {
        fmt.Println(i)
    }
    fmt.Println()

    i := 0
    for i <= 5 {
        fmt.Println(i)
        i++
    }
    fmt.Println()

    // rangeを使った繰り返し
    for i, v := range []int{10, 20, 30} {
        fmt.Println("i: ", i, "v: ", v)
    }
}
