package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	str1 := "this is a test str, first"
	str2 := "this is a test str, second"

	mSha1 := sha1.New()

	mSha1.Write([]byte(str1))
	ret1 := mSha1.Sum(nil)
	fmt.Printf("raw ret:%s, hex ret:%x \n", ret1, ret1)

	mSha1.Write([]byte(str2))
	ret2 := mSha1.Sum(nil)
	fmt.Printf("raw ret:%s, hex ret:%x \n", ret2, ret2)
}