package main

import "fmt"

func main() {
	fs := make([]func(), 3)
	for i := range fs {
		fmt.Println(i)
		fs[i] = func() { fmt.Println(i) }
	}

	println()
	fmt.Printf("%#v\n", fs)

	for _, f := range fs {
		f()
	}
}
