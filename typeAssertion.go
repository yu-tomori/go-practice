package main

import "fmt"

var v interface{}

func main() {
	v = 100
	n, ok := v.(int)
	fmt.Println(n, ok)

	s, ok := v.(string)
	fmt.Println(s, ok)
}
