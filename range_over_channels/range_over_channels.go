package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)

	queue := make(chan string, 5)
	queue <- "one"
	queue <- "two"
	queue <- "three"
	close(queue)

	go func() {
		defer wg.Done()
		for elem := range queue {
			fmt.Println(elem)
		}
	}()
	wg.Wait()
}

