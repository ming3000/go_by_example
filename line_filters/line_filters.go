package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
)

func main() {
	f, err := os.Open("_data/test_read_lines.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	myScanner := bufio.NewScanner(f)
	for myScanner.Scan() {
		lineText := myScanner.Text()
		fmt.Println(lineText)
		lineByte := myScanner.Bytes()
		fmt.Println(lineByte)
		fmt.Println(strings.Repeat("-", 50))
	}

	if err = myScanner.Err(); err != nil {
		panic(err)
	}
}
