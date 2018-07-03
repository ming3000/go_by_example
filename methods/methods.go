package main

import (
	"fmt"
	"strings"
)

type rect struct {
	width int
	height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2 * r.width + 2 * r.height
}

func main() {
	r := rect{width:10, height:5}
	fmt.Println("r area: ", r.area())
	fmt.Println("r perim: ", r.perim())
	fmt.Println(strings.Repeat("-", 10))

	rp := &r
	fmt.Println("rp area: ", rp.area())
	fmt.Println("rp perim: ", rp.perim())
	fmt.Println(strings.Repeat("-", 10))
}
