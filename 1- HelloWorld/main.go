package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")

	fmt.Println(addNumbers(5, 6))

	res1, res2 := addAndSubtractNumbers(5, 6)
	fmt.Println("output : ", res1, res2)

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	if val := 5; val < 10 {
		fmt.Println("ok")
	}
}

func addNumbers(x int, y int) int {
	return x + y
}

func addAndSubtractNumbers(x int, y int) (int, int) {
	return x + y, x - y
}
