package main

import (
	"fmt"
	"strings"
)

func main() {
	m := make(map[string]int)
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(m, n)
	fmt.Println(len(m), len(n))
	fmt.Println(strings.Repeat("-", 10))

	m["k1"] = 3
	m["k2"] = 7
	m["k3"] = 10
	fmt.Println("m:", m)
	fmt.Println("len m:", len(m))
	fmt.Println(strings.Repeat("-", 10))

	delete(m, "k2")
	fmt.Println("m:", m)
	fmt.Println("len m:", len(m))
	fmt.Println(strings.Repeat("-", 10))

	v, i := m["k3"]
	fmt.Println(v, i)
	v, i = m["k2"]
	fmt.Println(v, i)
	fmt.Println(strings.Repeat("-", 10))
}
