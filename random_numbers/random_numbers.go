package main

import (
	"math/rand"
	"fmt"
	"time"
	"strings"
)

func line() {
	fmt.Println(strings.Repeat("-", 30))
}

func main() {
	var randNum int

	randNum = rand.Int()
	fmt.Println(randNum)

	randNum = rand.Intn(100)
	fmt.Println(randNum)

	randNum = rand.Intn(10)
	fmt.Println(randNum)
	line()

	var randFloat float64
	randFloat = rand.Float64()
	fmt.Println(randFloat)
	randFloat = rand.Float64()
	fmt.Println(randFloat)
	line()

	src1 := rand.NewSource(42)
	gen1 := rand.New(src1)
	fmt.Println(gen1.Intn(10), gen1.Intn(10))

	// so the same source, generate the same rand num
	src2 := rand.NewSource(42)
	gen2 := rand.New(src2)
	fmt.Println(gen2.Intn(10), gen2.Intn(10))
	line()

	//use a variable source, will generate the variable num
	newSrc1 := rand.NewSource(time.Now().UnixNano())
	newGen1 := rand.New(newSrc1)
	fmt.Println(newGen1.Intn(10), newGen1.Intn(10))

	//use a variable source, will generate the variable num
	newSrc2 := rand.NewSource(time.Now().UnixNano())
	newGen2 := rand.New(newSrc2)
	fmt.Println(newGen2.Intn(10), newGen2.Intn(10))
	line()
}
