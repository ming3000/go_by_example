package main

import (
	"fmt"
	"strings"
)

type person struct {
	name string
	age int
}

func main() {
	var one person
	fmt.Println("one: ", one)

	one = person{name:"alice", age:30}
	fmt.Println("one: ", one)

	one = person{name:"fred"}
	fmt.Println("one: ", one)

	one = person{"jack", 20}
	fmt.Println("one: ", one)

	one = person{"pool", 35}
	fmt.Println("one.name: ", one.name)
	one.name = "dead pool"
	fmt.Println("one.name: ", one.name)

	fmt.Println(strings.Repeat("-", 10))

	two := &person{name:"ant", age:30}
	fmt.Println("two: ", two)
	fmt.Println("two.name: ", two.name)
	two.name = "ant man"
	fmt.Println("two.name: ", two.name)
}
