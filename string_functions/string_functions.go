package main

import (
	"fmt"
	"strings"
)

func line() {
	fmt.Println(strings.Repeat("-", 30))
}

func main() {
	fmt.Println("len:", len("hello"))
	line()

	fmt.Println("Char:", "hello"[1])
	line()

	fmt.Println("Count:", strings.Count("rabbitmq redis kafka", "r"))
	fmt.Println("Count:", strings.Count("rabbitmq redis kafka", "f"))
	line()

	fmt.Println("Contains:", strings.Contains("go in action", "ac"))
	fmt.Println("Index:", strings.Index("golang", "la"))
	line()

	fmt.Println("HasPrefix:", strings.HasPrefix("mongodb", "mon"))
	fmt.Println("HasSuffix:", strings.HasSuffix("redis", "os"))
	line()

	fmt.Println("Join:", strings.Join([]string{"a", "b", "c"}, "###"))
	fmt.Println("Split:", strings.Split("2018-06-22-17-40", "-"))
	line()

	fmt.Println("ToUpper:", strings.ToUpper("mysql"))
	fmt.Println("ToLower:", strings.ToLower("MongoDB"))
	line()

	fmt.Println(strings.Repeat("ha", 10))
	line()

	fmt.Println("Replace:", strings.Replace("fooobar", "o", "a", 1))
	// If n < 0, there is no limit on the number of replacements.
	fmt.Println("Replace:", strings.Replace("fooobar", "o", "a", -1))
	line()
}
