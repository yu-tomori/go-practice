package main

func swap(x, y int) (n, m int) {
	n, m = y, x
	return
}

func main() {
	n, m := swap(10, 20)
	println(n, m)
}
