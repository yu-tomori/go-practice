package main

func main() {
    p := struct {
        name string
        age int
    }{
        name: "Gopher",
        age: 10,
    }

    p.age++
    println(p.name, p.age)
}
