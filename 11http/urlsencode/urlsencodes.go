package main

import (
	"fmt"
	"net/url"
)

func main() {
	urls := []string{
		"https://example.com/hello world",
		"https://example.com/hello%20world",
		"https://example.com/你好",
		"https://example.com/%E4%BD%A0%E5%A5%BD",
	}

	for _, ustr := range urls {
		u, err := url.Parse(ustr)
		fmt.Println("======== Parsing URL:", ustr)
		if err != nil {
			fmt.Println("Url parse error!", err)
			continue
		}
		fmt.Println(u.RawPath)
		fmt.Println(u.EscapedPath())
	}
}
