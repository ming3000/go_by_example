package main

import (
	"sort"
	"fmt"
)

func main() {
	// 就地排序，改变原有序列，无返回值
	strs := []string{"e", "a", "g", "k"}
	sort.Strings(strs)
	fmt.Println(strs)

	ints := []int{98, 1, 33, 5, 76}
	sort.Ints(ints)
	fmt.Println(ints)
}
