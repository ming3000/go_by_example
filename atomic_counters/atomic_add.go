package main

import (
	"sync"
	"sync/atomic"
	"runtime"
	"fmt"
)

var counter int64
var wg sync.WaitGroup


func incCounter() {
	defer wg.Done()
	for count := 0; count < 2; count++ {
		atomic.AddInt64(&counter, 1)
		runtime.Gosched()
	}
}


func main() {
	wg.Add(2)

	go incCounter()
	go incCounter()

	wg.Wait()
	fmt.Println("final counter:", counter)
}
