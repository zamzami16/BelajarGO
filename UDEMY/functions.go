package main

import "fmt"

func Hellow(kali int, nama string) {
	for i := 0; i < kali; i++ {
		fmt.Println("Hello:", nama)
	}
}
func getPembagian(num float64, denum float64) float64 {
	return num / denum
}
func sumAll(numbers ...float64) (summed float64) {
	summed = 0
	for _, number := range numbers {
		summed += number
	}
	return summed
}

func main() {
	Hellow(10, "Zami")
	x := getPembagian(10, 1.3)
	fmt.Println(x)
	fmt.Println(sumAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 90))
}
