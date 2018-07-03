package main

import (
	"fmt"
	"strings"
)

func main() {
	a := make([]string, 5)
	b := []string{"go", "python", "javascript"}
	fmt.Println(a, b)
	fmt.Println(len(a), len(b))
	fmt.Println(strings.Repeat("-", 10))

	a[0] = "flask"
	a[1] = "gin"
	a[2] = "express"
	fmt.Println("a:", a)
	fmt.Println(strings.Repeat("-", 10))

	// notice the len, a[3], a[4] is default value
	a = append(a, "django")
	fmt.Println("a:", a)
	fmt.Println("len a:", len(a))
	fmt.Println(strings.Repeat("-", 10))

	a = append(a, "yii", "laravel")
	fmt.Println("a:", a)
	fmt.Println("len a:", len(a))
	fmt.Println(strings.Repeat("-", 10))

	c := make([]string, len(a))
	copy(c, a)
	fmt.Println("copy a:", c)
	fmt.Println(strings.Repeat("-", 10))

	l := a[3:5]
	fmt.Println("notice, l, 3:5", l)
	l = a[3:6]
	fmt.Println("l, 3:6", l)
	l = a[:6]
	fmt.Println("l, :6", l)
	l = a[5:]
	fmt.Println("l, 5:", l)
	fmt.Println(strings.Repeat("-", 10))
}

