package main

import (
	"fmt"
)

var i interface{}

func main() {
	i = 100
	switch v := i.(type) {
	case int:
		fmt.Println(v * 2)
	case string:
		fmt.Println(v + "hoge")
	default:
		fmt.Println("default")
	}
}
