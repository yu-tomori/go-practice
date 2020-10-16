package main

func main() {
    ns := make([]int, 3, 10)
    // 上と下はだいたい同じ処理
    var array [10]int
    ns := array[0:3] // またはarray[:3]

    ms := []int{10, 20, 30, 40, 50}
    // 上と下はだいたい同じ処理
    var array2 = [...]int{10, 20, 30, 40, 50}
    ms := array2[0:5]
}
