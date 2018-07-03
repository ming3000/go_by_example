package main

import "fmt"

const g_s string = "the global constant"

func main() {
	fmt.Println(g_s)

	const l_s = "the local constant"
	fmt.Println(l_s)
}
