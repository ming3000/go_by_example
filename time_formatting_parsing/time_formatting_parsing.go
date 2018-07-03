package main

import (
	"fmt"
	"strings"
	"time"
	"reflect"
)

func line() {
	fmt.Println(strings.Repeat("-", 30))
}

func main() {
	t := time.Now()
	tStr := fmt.Sprintf("string_datetime:%d_%02d_%02d__%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	fmt.Println(tStr)
	line()

	fmt.Printf("current_date:%d_%02d_%02d \n",
		t.Year(), t.Month(), t.Day())
	fmt.Printf("current_time:%02d:%02d:%02d \n",
		t.Hour(), t.Minute(), t.Second())
	fmt.Printf("current_datetime:%d_%02d_%02d__%02d:%02d:%02d \n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	line()

	fmt.Println(t.Format(time.ANSIC))
	fmt.Println(t.Format(time.RFC3339))
	line()

	// second must be 0x format, can not be xx format
	fmt.Println(t.Format("2006-01-02 15:30:02"))
	fmt.Println(t.Format("2006-01-02 15:30:20"))
	fmt.Println(t.Format("2006-01-02T15:03"))
	line()

	pt, err := time.Parse(time.RFC3339, "2018-05-03T09:10:30+08:00")
	if err != nil {
		panic(err)
	}
	fmt.Println(reflect.TypeOf(pt), pt)
	line()

	// the layout must be 2006-01-02 15:04:05, others wont be affected
	pt, err = time.Parse("2006-01-02 15:04:05", "2018-05-03 09:18:30")
	if err != nil {
		panic(err)
	}
	fmt.Println(reflect.TypeOf(pt), pt)
	line()

	pt, err = time.Parse("2006-01-02", "2018-06-23")
	if err != nil {
		panic(err)
	}
	fmt.Println(reflect.TypeOf(pt), pt)
	line()

	// anyway, the layout must be the day of "2006 01 02"
	pt, err = time.Parse("2006_01_02", "2018_06_23")
	if err != nil {
		panic(err)
	}
	fmt.Println(reflect.TypeOf(pt), pt)
	line()

	pt, err = time.Parse("2006_01_02_15_04_05", "2018_06_23_09_30_02")
	if err != nil {
		panic(err)
	}
	fmt.Println(reflect.TypeOf(pt), pt)
	line()
}
