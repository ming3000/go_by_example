package main

import (
	"fmt"
	"os"
)

func testDefer() {
	defer fmt.Println("can you see me")
}

func main() {
	testDefer()

	// won't be print
	defer fmt.Println("you can't see me")
	os.Exit(233)
}

