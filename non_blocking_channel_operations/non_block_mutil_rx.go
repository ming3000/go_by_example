package main

import (
	"time"
	"fmt"
)

func main() {
	someChan := make(chan string)
	otherChan := make(chan string)

	go func() {
		time.Sleep(time.Second * 3)
		someChan <- "some text"
	}()

	go func() {
		time.Sleep(time.Second * 5)
		otherChan <- "other text"
	}()

	select {
	case res1 := <-someChan:
		fmt.Println("some chan,", res1)
	case res2 := <-otherChan:
		fmt.Println("other chan,", res2)
	default:
		fmt.Println("default, no block")
	}
}
