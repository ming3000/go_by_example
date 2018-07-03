package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
	"bufio"
)

func line() {
	fmt.Println(strings.Repeat("-", 30))
}

func main() {
	data1 := []byte("hello golang\n")
	err := ioutil.WriteFile("_data/test_write_1.txt",
							data1, 0644)
	if err != nil {
		panic(err)
	}

	f, err := os.Create("_data/test_write_2.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	data2 := []byte("my blood gets heavy when you are here...\n")
	count2, _ := f.Write(data2)
	fmt.Printf("write %d bytes \n", count2)

	// notice the \n here
	count3, _ := f.WriteString("i can't turn my soul, when you are here...\n")
	fmt.Printf("write %d bytes \n", count3)

	// sync
	f.Sync()

	bufFile := bufio.NewWriter(f)
	count4, _ := bufFile.WriteString("everything gone...\n")
	fmt.Printf("write %d bytes \n", count4)
	// flush
	bufFile.Flush()
}

