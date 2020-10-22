package main

import "fmt"

func main() {
	ns := [3]int{10, 20, 30}
	ns2 := &ns
	fmt.Println("ns: ", ns)
	fmt.Println("ns2: ", ns2)

	ns[1] = 200
	fmt.Println("ns: ", ns)
	fmt.Println("ns2: ", ns2)

	ns2[1] = 100
	fmt.Println("ns: ", ns)
	fmt.Println("ns2: ", ns2)
}
