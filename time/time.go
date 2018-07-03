package main

import (
	"time"
	"fmt"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Year(), now.Month(), now.Day())
	fmt.Println(now.Hour(), now.Minute(), now.Second())

	fmt.Println(now.Location())
	fmt.Println(now.Weekday())

	sometime := time.Date(2009, 7, 20,
					20, 30, 50,
					651387237,
					time.UTC)
	fmt.Println(sometime)

	tomorrow := now.Add(24 * time.Hour + 10 * time.Minute + 30 * time.Second)
	fmt.Println(tomorrow)
	fmt.Println(tomorrow.Before(now))
	fmt.Println(tomorrow.After(now))
	fmt.Println(tomorrow.Equal(now))

	diff := tomorrow.Sub(now)
	fmt.Println(diff)
	fmt.Println(diff.Hours(), diff.Minutes(), diff.Seconds())

	// Duration is the same as int64
	var dela time.Duration = 60 * 60 * time.Second
	next := now.Add(dela)
	fmt.Println(next)
}
