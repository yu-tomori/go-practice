package main
import "fmt"

func main() {
    var i int
    for {
        if i  >= 5 {
            break
        }
        fmt.Println(i)
        i++
    }
    fmt.Println()

LOOP:
    for {
        switch {
        case i % 2 == 1:
            break LOOP
        }
        i++
    }
}
