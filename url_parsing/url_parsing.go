package main

import (
	"net/url"
	"fmt"
	"strings"
)

func line() {
	fmt.Println(strings.Repeat("-", 30))
}

func myParser(urlStr string) {
	urlDst, err := url.Parse(urlStr)
	if err != nil {
		panic(err)
	}

	fmt.Println(urlDst.Scheme)
	fmt.Println(urlDst.Host)
	fmt.Println(urlDst.Path)
	fmt.Println(urlDst.Fragment)
	line()

	fmt.Println(urlDst.User)
	fmt.Println(urlDst.User.Username())
	fmt.Println(urlDst.User.Password())
	line()

	fmt.Println(urlDst.RawQuery)
	m, _ := url.ParseQuery(urlDst.RawQuery)
	fmt.Println(m)
}



func main() {
	pgs := "postgres://user:pass@host.com:5432/path?k=v#f"
	myParser(pgs)
	fmt.Println(strings.Repeat(">", 50))
	mgs := "mongodb://db1.example.net:27017,db2.example.net:2500/test_db/?replicaSet=test"
	myParser(mgs)
	fmt.Println(strings.Repeat(">", 50))
	mons := "mongodb://admin:123456@localhost/test_db/?_id=111"
	myParser(mons)
}
