package main

import (
	"sync"
	"fmt"
)

var wGroup sync.WaitGroup

func init() {
	wGroup.Add(1)
}

func main() {
	msgs := make(chan string)
	go testChannel(msgs)

	back := <- msgs
	fmt.Println("rx message:", back)

	wGroup.Wait()
	fmt.Println("main meet the end")
}

func testChannel(ch chan string) {
	fmt.Println("a new routine start")
	defer wGroup.Done()
	ch <- "ping ping ping"
}


