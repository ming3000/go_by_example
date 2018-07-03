package main

import (
	"fmt"
	"strings"
)

func intSeq() (func() int) {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	nextInt := intSeq()
	fmt.Println("the next int:", nextInt())
	fmt.Println("the next int:", nextInt())
	fmt.Println("the next int:", nextInt())
	fmt.Println(strings.Repeat("-", 10))

	newInt := intSeq()
	fmt.Println("the new int:", newInt())
	fmt.Println("the new int:", newInt())
	fmt.Println("the new int:", newInt())
	fmt.Println(strings.Repeat("-", 10))
}
