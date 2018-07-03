package main

import (
	"time"
	"fmt"
)

func main() {
	msg := make(chan string)
	go func() {
		time.Sleep(time.Second * 5)
		msg <- "head"
	}()

	select {
	case res := <-msg:
		fmt.Println("main rx: ", res)
	default:
		fmt.Println("default, no block")
	}
}
