package main

import "fmt"

/**
Interface a/ tipe data abstract, tidka punya implementasi langsung
*/

type Person struct {
	Name string
}

type HasName interface {
	GetName() string
}

func SayHello(hasName HasName) {
	fmt.Println("Hello ", hasName.GetName())
}

func (person Person) GetName() string {
	return person.Name
}

func main() {
	aldi := Person{Name: "Aldi Taher"}

	SayHello(aldi)
}
