package main

import "fmt"

type T int

func (t T) M() {
	fmt.Print(int(t))
}

func main() {
	ts := []T{10, 20, 30}

	ms := make([]func(), 0, len(ts))
	for _, t := range ts {
		// tはforの中でずっと同じポインタ
		ms = append(ms, t.M)
	}

	// 303030と表示される
	for _, m := range ms {
		m()
	}
}
