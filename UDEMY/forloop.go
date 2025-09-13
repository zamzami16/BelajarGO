package main

import "fmt"

func main() {
	idx := 0
	for idx < 10 {
		fmt.Println(idx)
		idx++
	}

	for idx := 1; idx < 200; idx++ {
		fmt.Println("jancok ", idx)
	}

	slice := []string{"budi", "joko", "santoso", "mangan", "telo"}
	for idx, name := range slice {
		fmt.Println(idx, name)
	}

	maps := map[string]string{
		"name": "zami",
		"age":  "20",
	}
	for key, val := range maps {
		fmt.Println(key, "=", val)
	}

	maps1 := map[int]string{
		1: "eko",
		2: "kurniawan",
		3: "Naniiii",
	}
	for key, val := range maps1 {
		fmt.Println(key, val)
	}
}
