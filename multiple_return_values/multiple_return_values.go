package main

import "fmt"

func vals(a int, b int) (int, int) {
	return a, b
}

func main() {
	a, b := vals(3, 5)
	fmt.Println("return value:", a, b)

	_, c := vals(7, 11)
	fmt.Println("return value:", c)
}
