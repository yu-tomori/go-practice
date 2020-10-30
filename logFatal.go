package main

import (
	"log"
)

func f() {
	var n int = 2
	println(n)
}

func main() {
	if err := f(); err != nil {
		log.Fatal(err)
	}
}
