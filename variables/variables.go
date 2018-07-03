package main

import (
	"fmt"
	"strings"
)

func main() {
	var a string = "initial"
	var b int = 3
	fmt.Println(a, b)
	fmt.Println(strings.Repeat("-", 30))

	var aa string
	var bb int
	fmt.Println(aa, bb)
	fmt.Println(strings.Repeat("-", 30))

	var c, d int = 1, 2
	fmt.Println(c, d)
	var cc, dd = 33, "ss"
	fmt.Println(cc, dd)
	fmt.Println(strings.Repeat("-", 30))

	var e = true
	var f = 3.14
	fmt.Println(e, f)
	fmt.Println(strings.Repeat("-", 30))

	x := "short"
	y := false
	fmt.Println(x, y)
	fmt.Println(strings.Repeat("-", 30))
}
