package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) != 3 {
		println("input error")
		os.Exit(1)
	}

	num1, _ := strconv.Atoi(os.Args[1])
	num2, _ := strconv.Atoi(os.Args[2])

	size := math(num1, num2)

	for x := 0; x < size; x++ {
		fmt.Println("Hello, World!")
	}
}

func math(num1 int, num2 int) int {
	result := num1 + num2
	return result
}
