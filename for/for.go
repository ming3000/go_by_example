package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	i := 1
	for i < 3 {
		fmt.Println(i)
		i++
	}

	for j := 5; j < 10; j++ {
		fmt.Println(j)
	}

	inputReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("typein >>>")
		in, _ := inputReader.ReadString('\n')
		in = strings.Trim(in, "\n")
		if in == "exit" {
			break
		}
		fmt.Println("input: ", in)
	}
}
