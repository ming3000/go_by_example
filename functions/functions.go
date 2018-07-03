package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func main()  {
	ret := plus(3, 5)
	fmt.Println("result of plus is ", ret)
}
