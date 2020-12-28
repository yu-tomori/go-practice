package main

import (
	"crypto/sha256"
	"fmt"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	var count int = CompareSHA(c1, c2)
	fmt.Println(count)
}

func CompareSHA(c1, c2 [32]byte) int {
	var count int = 0
	for i := 0; i < len(c1); i++ {
		count += int(pc[c1[i]&c2[i]])
		// fmt.Printf("%b %b %b %v\n", c1[i], c2[i], c1[i]&c2[i], int(pc[c1[i]&c2[i]]))
	}
	return count
}
