package main

func swap(x, y *int) {
	*x, *y = *y, *x
}

func main() {
	n, m := 10, 20
	swap(&n, &m)
	println(n, m)
}
