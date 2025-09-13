package main

import "fmt"

func main() {
	var newSlice = make([]string, 2, 4)
	newSlice[0] = "ayam"
	newSlice[1] = "pitek"
	fmt.Println(newSlice)
	fmt.Println(cap(newSlice))
	fmt.Println(len(newSlice))
}
