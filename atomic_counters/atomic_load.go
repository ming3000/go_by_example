package main

import (
	"sync/atomic"
	"runtime"
	"fmt"
	"time"
)

func main() {
	var ops uint64 = 0

	for i := 0; i < 1000; i++ {
		go func() {
			atomic.AddUint64(&ops, 1)
			runtime.Gosched()
		}()
	}

	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)

	time.Sleep(time.Second)
	opsFinal = atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}
