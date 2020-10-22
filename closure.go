package main

func main() {
	msg := "Hello, 世界"
	func() {
		println(msg)
	}
}
