package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, rets chan<- string) {
	for j := range jobs {
		fmt.Println("worker:", id, "processing job:", j)
		time.Sleep(time.Second)
		rets <- fmt.Sprintf("result of the job %d", j)
	}
}

func main() {
	jobs := make(chan int, 100)
	rets := make(chan string, 100)

	for w :=0; w < 3; w++ {
		go worker(w, jobs, rets)
	}

	for j := 0; j < 10; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 0; a < 10; a++ {
		ret := <- rets
		fmt.Println(ret)
	}
}
