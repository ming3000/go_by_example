package main

import (
	"time"
	"fmt"
)

func main() {
	ch1 := make(chan string)
	go func() {
		time.Sleep(time.Second * 5)
		ch1 <- "result 1"
	}()

	select {
	case res1 := <-ch1:
		fmt.Println("rx: ", res1)
	case <-time.After(time.Second * 2):
		fmt.Println("2 second, timeout...")
	}

	// -------------------------
	ch2 := make(chan string)
	go func () {
		time.Sleep(time.Second * 2)
		ch2 <- "result 2"
	}()
	select {
	case res2 := <- ch2:
		fmt.Println("rx: ", res2)
	case <-time.After(time.Second * 5):
		fmt.Println("5 second, timeout...")
	}
}
