package main
import (
    "fmt"
    "time"
    "math/rand"
)

func main() {
    rand.Seed(time.Now().UnixNano())
    result := rand.Intn(6)
    result++
    switch result {
    case 6:
        fmt.Println("大吉")
    case 4, 5:
        fmt.Println("中吉")
    case 2, 3:
        fmt.Println("吉")
    default:
        fmt.Println("凶")
    }
}
