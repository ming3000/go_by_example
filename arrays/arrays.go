package main

import (
	"fmt"
	"strings"
)

func main() {
	var a [5]int
	var b [3]bool
	fmt.Println("array a:", a)
	fmt.Println("array b:", b)
	fmt.Println(strings.Repeat("-", 10))

	a[3] = 3
	fmt.Println("array a all:", a)
	fmt.Println("array a one:", a[3])
	fmt.Println(strings.Repeat("-", 10))

	fmt.Println("array a len:", len(a))
	fmt.Println("array b len:", len(b))
	fmt.Println(strings.Repeat("-", 10))

	var c [3]string = [3]string{"python", "go", "javascript"}
	fmt.Println("array c:", c)
	fmt.Println(strings.Repeat("-", 10))
}
