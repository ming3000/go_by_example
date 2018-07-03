package main

import (
	"fmt"
	"time"
)

func main() {
	flagChan := make(chan bool)
	go workerRoutine(flagChan)

	ret := <-flagChan
	if ret {
		fmt.Println("work done")
	} else {
		fmt.Println("work delay")
	}
	fmt.Println("main meet the end")
}

func workerRoutine(ch chan bool) {
	fmt.Println("worker routine start working...")
	time.Sleep(time.Second * 5)
	fmt.Println("worker routine done work...")
	ch <- true
}
