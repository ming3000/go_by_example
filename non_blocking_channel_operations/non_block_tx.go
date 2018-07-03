package main

import (
	"fmt"
	"time"
)

func main() {
	msg := make(chan string, 1)

	go func() {
		time.Sleep(time.Second * 3)
		fmt.Println("routine,", <-msg)
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg <- "hahaha":
			fmt.Printf("main, the %d times tx\n", i)
		default:
			fmt.Println("default, no block")
		}
	}
}
