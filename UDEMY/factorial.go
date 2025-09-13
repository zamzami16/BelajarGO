package main

import "fmt"

func factorialLoop(value int) int {
	result := 1
	if value == 0 {
		return result
	}
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

func factorialRecursive(value int) int {
	if value == 1 || value == 0 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func main() {
	x := 10
	loop := factorialLoop(x)
	loopRec := factorialRecursive(x)
	fmt.Println(loop, loopRec)
}
