package main

import (
	"time"
	"fmt"
)

func main() {
	now := time.Now()
	fmt.Println(now)

	secs := now.Unix()
	fmt.Println(secs)

	day := time.Unix(secs, 0)
	fmt.Println(day)
}
