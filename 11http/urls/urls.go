package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	urlstr := "http://testuse:testpd@www.google.com:9090/test/page1?a=b&3=8#frag23"

	// /test/page#frag23?a=b&3=8 , #frag23?a=b&3=8 会被认为是fragment，此时拿不到rawQuery
	fmt.Println(urlstr)
	u, err := url.Parse(urlstr)
	if err != nil {
		panic(err)
	}
	fmt.Println(u)
	fmt.Println(u.String())
	fmt.Println(u.Scheme)
	fmt.Println(u.User.String())
	fmt.Println(u.User.Username())
	fmt.Println(u.User.Password())
	fmt.Println(u.Host)
	fmt.Println(u.Path)
	fmt.Println(u.RawQuery)
	fmt.Println(u.Fragment)

	host, p, err := net.SplitHostPort(u.Host)
	if err != nil {
		fmt.Println("No Host in url", err)
	}
	fmt.Println(host)
	fmt.Println(p)

	fmt.Println(u.Query())

	// 直接使用u.Query()  无需使用url.ParseQuery了
	query, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		panic(err)
	}

	fmt.Println(query)
	fmt.Println(query.Get("a"))

}
