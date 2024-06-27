package main

import "fmt"

func main() {
	p := Person{name: "John"}

	changeName := func(newName string) {
		p.name = newName
	}

	changeName("Jane")

	fmt.Println(p.name)
}
