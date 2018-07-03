package main

import (
	"fmt"
	"sync"
)

var wGroup sync.WaitGroup

func init() {
	wGroup.Add(1)
}

func main() {
	msgs := make(chan string, 2)
	go testRoutine(msgs)

	fmt.Println("start rx")
	rx1 := <-msgs
	rx2 := <-msgs
	fmt.Println("after rx 1:", rx1)
	fmt.Println("after rx 2:", rx2)

	wGroup.Wait()
	fmt.Println("main meet the end")
}

func testRoutine(ch chan string) {
	defer wGroup.Done()
	fmt.Println("a new routine start")

	ch <- "ping ping"
	ch <- "bang bang"
}
