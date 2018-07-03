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
	data, err := ioutil.ReadFile("_data/test_read.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
	fmt.Println(strings.Repeat(">", 50))

	f, err := os.Open("_data/test_read.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	byte1 := make([]byte, 5)
	count1, _ := f.Read(byte1)
	fmt.Printf("read: %d bytes, strings is: %s \n", count1, string(byte1))
	fmt.Println(strings.Repeat(">", 50))

	offset2, _ := f.Seek(6, 0)
	byte2 := make([]byte, 10)
	count2, _ := f.Read(byte2)
	fmt.Printf("read %d bytes @ %d: %s \n", count2, offset2, byte2)

	// reset the offset pointer
	offset3, _ := f.Seek(0, 0)
	byte3 := make([]byte, 10)
	count3, _ := f.Read(byte3)
	fmt.Printf("read %d bytes @ %d: %s \n", count3, offset3, byte3)
	fmt.Println(strings.Repeat(">", 50))

	// buf io
	bufedFile := bufio.NewReader(f)
	byte4, _ := bufedFile.Peek(10)
	fmt.Println("bufed read 10, byte:", byte4)
	fmt.Println("bufed read 10, string:", string(byte4))
	fmt.Println(strings.Repeat(">", 50))
}
