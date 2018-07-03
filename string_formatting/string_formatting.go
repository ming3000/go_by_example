package main

import (
	"fmt"
	"strings"
	"os"
)

type point struct {
	x int
	y int
}

func strFormat() {
	fmt.Println(strings.Repeat("-", 30))
	s := fmt.Sprintf("you like %s, i like %s.", "java", "golang")
	fmt.Println(s)
	fmt.Println(strings.Repeat("-", 30))
}

func fileFormat() {
	fmt.Println(strings.Repeat("-", 30))
	fmt.Fprintf(os.Stdout, "aways %s sky! \n", "blue")

	f, _ := os.Create("_data/string_formatting")
	defer f.Close()
	fmt.Fprintf(f, "aways %s sky! \n", "blue")

	fmt.Println(strings.Repeat("-", 30))
}

func main() {
	p := point{3, 7}
	fmt.Printf("%v \n", p)
	fmt.Printf("%+v \n", p)
	fmt.Printf("%#v \n", p)

	fmt.Printf("%T \n", p)
	fmt.Printf("%p \n", &p)

	fmt.Printf("the bool %t \n", true)
	fmt.Printf("the bool %t \n", false)

	fmt.Printf("%d \n", 123)
	fmt.Printf("%b \n", 7)
	fmt.Printf("%x \n", 456)
	fmt.Printf("%c \n", 33)

	fmt.Printf("%f \n", 78.9)

	fmt.Printf("the %s \n", "way you lie")

	strFormat()
	fileFormat()
}
