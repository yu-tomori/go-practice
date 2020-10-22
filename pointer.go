package main

func f(xp *int) { // the type 'pointer' of int
	*xp = 100 // insert the value to which the pointer refering
}

func main() {
	var x int
	f(&x) // get the pointer by add prefix '&'
	println(x)
}
