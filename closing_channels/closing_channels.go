package main

import (
	"sync"
	"fmt"
)

var wGroup sync.WaitGroup

func main() {
	wGroup.Add(1)
	jobs := make(chan int, 5)

	go func() {
		defer wGroup.Done()
		for {
			i, more := <-jobs
			if more {
				fmt.Println("rx job:", i)
			} else {
				fmt.Println("all job received")
				return
			}
		}
	}()

	for j := 0; j < 3; j++ {
		jobs <- j
	}
	close(jobs)
	wGroup.Wait()
}
