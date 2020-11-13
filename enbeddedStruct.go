package main

import "fmt"

type Name struct {
	FirstName string
	LastName  string
}

func (n *Name) String() string {
	return fmt.Sprintf("%s, %s", n.FirstName, n.LastName)
}

type Person struct {
	// *Name型の値を埋め込む
	*Name
	Age int
}

func main() {
	n := &Name{
		FirstName: "Taro",
		LastName:  "Yamada",
	}

	p := &Person{
		Name: n,
		Age:  20,
	}
	fmt.Println(p.String())

	fmt.Printf("type of p is %T\n", p)
	var stringer fmt.Stringer
	fmt.Printf("[Before insertion] type of stringer is %T\n", stringer)
	stringer = p
	fmt.Printf("[After insertion] type of stringer is %T\n", stringer)

	println()
	var i interface{}
	fmt.Printf("the type of i is %T\n", i)
}
