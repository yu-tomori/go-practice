package main

func main() {
    a := []int{19, 86, 1, 12}
    var sum int = 0
    for _, v := range a {
        sum += v
    }
    println(sum)
}
