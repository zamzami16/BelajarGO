package main

import "fmt"

// People Buat struct
type People struct {
	Name, Address string
	Age           int
}

// Buat struct method
func (people People) sayHi() {
	fmt.Println("Hi, my name is", people.Name)
}

func main() {
	dayah := People{
		Name:    "Dayah",
		Address: "Mangunan",
		Age:     25,
	}
	fmt.Println(dayah)
	dayah.sayHi()
}
