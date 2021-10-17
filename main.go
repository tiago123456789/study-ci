package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func multiple(a int, b int) int {
	return a * b
}

func divise(a int, b int) int {
	return a / b
}

func main() {
	fmt.Printf("THe value sum method => %d \n", sum(10, 10))
}
