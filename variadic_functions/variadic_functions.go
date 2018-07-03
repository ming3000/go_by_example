package main

import (
	"fmt"
	"strings"
)

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println("input:", nums)
	fmt.Println("result:", total)
	fmt.Println(strings.Repeat("-", 10))
	return total
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4, 5}
	sum(nums...)
}
