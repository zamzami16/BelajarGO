package main

import (
	"fmt"
	"strconv"
)

func main() {
	const (
		ayam  = "cemani"
		bento = "koen"
	)
	name := "Zami"
	age := 20
	ages := "20"
	fmt.Println(name, age, ayam, bento)
	i, err := strconv.Atoi(ages)
	if err != nil {
		panic(err)
	}
	fmt.Println(i)
}
