package main

import (
	"fmt"
)

func main() {
	arrayNums := [3]int{1, 2, 3}
	for _, num := range arrayNums {
		fmt.Print(num, ",")
	}
	fmt.Println()

	sliceNums := []int{5, 6, 7}
	for i, num := range sliceNums {
		fmt.Printf("index %d is %d, ", i, num)
	}
	fmt.Println()

	mapNums := map[string]string{"a": "python", "b": "golang", "c": "node.js"}
	for k, v := range mapNums {
		fmt.Printf("key %s is %s, ", k, v)
	}
	fmt.Println()

	for i, c := range "node.js" {
		fmt.Println(i, c)
	}
}

