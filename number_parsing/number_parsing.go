package main

import (
	"fmt"
	"strings"
	"strconv"
	"reflect"
)

func line() {
	fmt.Println(strings.Repeat("-", 30))
}

func main() {
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(reflect.TypeOf(i), i)

	// bitsize useless
	i, _ = strconv.ParseInt("567",0, 32)
	fmt.Println(reflect.TypeOf(i), i)
	line()

	// base is 8, 10, 16, 0 is auto detect
	i, _ = strconv.ParseInt("0xFF",0, 64)
	fmt.Println(reflect.TypeOf(i), i)

	i, _ = strconv.ParseInt("123", 16, 64)
	fmt.Println(reflect.TypeOf(i), i)

	i, _ = strconv.ParseInt("0xFF", 10, 64)
	fmt.Println(reflect.TypeOf(i), i)
	line()

	// bitsize is 32, 64
	f, _ := strconv.ParseFloat("3.14", 64)
	fmt.Println(reflect.TypeOf(f), f)

	b, _ := strconv.ParseBool("false")
	fmt.Println(reflect.TypeOf(b), b)
	line()

	// shortcut
	v, e := strconv.Atoi("999")
	fmt.Println(v, e)

	v, e = strconv.Atoi("diango")
	fmt.Println(v, e)
	line()
}
