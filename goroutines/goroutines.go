package main

import (
	"sync"
	"fmt"
	"strings"
)

var wGroup sync.WaitGroup

func init() {
	wGroup.Add(2)
}

func main() {
	go testRoutine("test routine 1")
	go testRoutine("test routine 2")

	wGroup.Wait()
	fmt.Println(strings.Repeat(">", 50))
	fmt.Println("main meet the end")
}

func testRoutine(from string) {
	defer wGroup.Done()

	for i := 0; i < 10; i++ {
		fmt.Println(from, ": ", i)
	}
}
