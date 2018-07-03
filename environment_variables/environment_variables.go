package main

import (
	"os"
	"fmt"
	"strings"
)

func main() {
	os.Setenv("foo", "233")

	ret := os.Getenv("foo")
	fmt.Println("env foo:", ret)

	ret = os.Getenv("bar")
	fmt.Println("env bar:", ret)
	fmt.Println(strings.Repeat(">", 50))

	for _, value := range os.Environ() {
		valuePair := strings.Split(value, "=")
		fmt.Println("env:", valuePair[0])
	}
	fmt.Println(strings.Repeat(">", 50))
}
