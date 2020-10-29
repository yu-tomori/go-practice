package main

import (
	"fmt"
	greeting "github.com/tenntenn/greeting/v2"
	mysync "github.com/tenntenn/sync"
	"sync"
)

func main() {
	fmt.Println(greeting.Do())
}
