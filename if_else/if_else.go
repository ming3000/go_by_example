package main

import "fmt"

func main() {
	flag := false
	if flag {
		fmt.Println("is true")
	} else {
		fmt.Println("is false")
	}

	name := "golang"
	if name == "golang" {
		fmt.Println("name is", name)
	}

	if num := 1; num < 0 {
		fmt.Println(num, " is small than 0")
	} else {
		fmt.Println(num, " is big than 0")
	}
}


