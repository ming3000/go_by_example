package main

import (
	"os"
	"fmt"
)

func createFile(p string) *os.File {
	fmt.Println("> createFile")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println(">> writeFile")
	fmt.Fprintln(f, "defer demo comtent")
}

func closeFile(f *os.File) {
	fmt.Println(">>> closeFile")
	f.Close()
}

func main() {
	f := createFile("_data/test_write3.txt")
	defer closeFile(f)

	writeFile(f)
}
