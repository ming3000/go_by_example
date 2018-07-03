package main

import (
	"time"
	"fmt"
)

// select,如果有default子句，则执行该子句
// 如果没有default子句，select将阻塞，直到某个通道返回
// 如果有多个case可以运行，select会随机选出一个执行。其他不会执行
func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 5)
		c1 <- "first chan"
	}()

	go func() {
		time.Sleep(time.Second * 1)
		c2 <- "second chan"
	}()

	select {
	case msg1 := <- c1:
		fmt.Println("rx: ", msg1)
	case msg2 := <- c2:
		fmt.Println("rx: ", msg2)
	}
}
