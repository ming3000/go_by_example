package main

import "fmt"

func zeroValue(iValue int) {
	iValue = 0
}

func zeroPtr(iPtr *int) {
	*iPtr = 0
}


func main() {
	i := 1
	fmt.Println("initial, value:", i)
	fmt.Println("pointer, value:", &i)

	zeroValue(i)
	fmt.Println("after zeroValue, value:", i)

	zeroPtr(&i)
	fmt.Println("after zeroPtr, value:", i)
}
