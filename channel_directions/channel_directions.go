package main

import (
	"fmt"
	"time"
)

func badEcho(in <-chan string) {
	// invalid operation: in <- "bad bad",
	// send to receive-only type <-chan string
	in <- "bad bad"
}

func testEcho(in <-chan string, out chan<- string) {
	inStr := <-in
	fmt.Println("routine rx:", inStr)
	out <- inStr
}

func main() {
	inChan := make(chan string)
	outChan := make(chan string)
	go testEcho(outChan, inChan)

	fmt.Println("main tx a message")
	outChan <- "All of the dust and dirt in the ground at some point"

	time.Sleep(time.Second * 2)
	backMsg := <- inChan
	fmt.Println("main rx a message: " + backMsg)

	badChan := make(chan string)
	go badEcho(badChan)
}
