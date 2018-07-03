package main

import (
	"sync"
	"fmt"
	"time"
)

func main() {
	songs := make([]string, 0)
	wg := sync.WaitGroup{}
	var mutex = &sync.Mutex{}

	wg.Add(30)
	for i := 0; i < 30; i++ {
		go func(id int) {
			defer wg.Done()
			mutex.Lock()
			ret := fmt.Sprintf("routine %d enter,", id)
			songs = append(songs, ret)
			fmt.Println(ret)
			time.Sleep(time.Second)
			mutex.Unlock()
		}(i)
	}

	wg.Wait()
	mutex.Lock()
	fmt.Println("main enter")
	fmt.Println(songs)
	mutex.Unlock()
}
