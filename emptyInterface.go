package main

import "fmt"

func main() {
	var v interface{}
	v = 100
	v = "hoge"
	fmt.Println(v)
}
