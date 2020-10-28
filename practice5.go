package main

func isEven(x int) (y bool) {
	if x%2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	for i := 0; i <= 100; i++ {
		print(i)
		if isEven(i) {
			println("-偶数")
		} else {
			println("-奇数")
		}
	}
}
