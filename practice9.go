package main

import "fmt"

type Stringer interface {
	String() string
}

type Human struct {
	Name    string
	Parent  string
	Food    string
	History int
}

func (h Human) String() string {
	return fmt.Sprintf("%v Parent: %v, Food: %v (%d years)", h.Name, h.Parent, h.Food, h.History)
}

type Animal struct {
	Name    string
	Parent  string
	Food    string
	History int
}

func (a Animal) String() string {
	return fmt.Sprintf("%v Parent: %v, Food: %v (%d years)", a.Name, a.Parent, a.Food, a.History)
}

type Plant struct {
	Name    string
	Parent  string
	Food    string
	History int
}

func (p Plant) String() string {
	return fmt.Sprintf("%v Parent: %v, Food: %v (%d years)", p.Name, p.Parent, p.Food, p.History)
}

func typeSwitch(stri interface{}) {
	switch stri.(type) {
	case Plant:
		fmt.Printf("Plant: ")
		fmt.Println(stri)
	case Human:
		fmt.Printf("Human: ")
		fmt.Println(stri)
	case Animal:
		fmt.Printf("Animal: ")
		fmt.Println(stri)
	default:
		fmt.Printf("Unknown: ")
		fmt.Println(stri)
	}
}

func main() {
	var flower Stringer = Plant{
		Name:    "Flower",
		Parent:  "Seas",
		Food:    "Water",
		History: 10000000,
	}
	var tiger Stringer = Animal{
		Name:    "Tiger",
		Parent:  "kapibara",
		Food:    "meat",
		History: 100000,
	}
	var yusuke Stringer = Human{
		Name:    "Yusuke",
		Parent:  "Ayumi",
		Food:    "gyoza",
		History: 21,
	}

	typeSwitch(flower)
	typeSwitch(tiger)
	typeSwitch(yusuke)

}
