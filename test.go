package main

import "fmt"

var s []int = []int{1, 2, 3, 4}

func main() {
	lice := s // assing by pointer
	s[0] = 0
	fmt.Printf("s: %v\n", s)       // Output: s: [0 2 3 4]
	fmt.Printf("lice: %v\n", lice) // Output: lice: [0 2 3 4]
}
