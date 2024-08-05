package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano() / 1000000)
	fmt.Println(now.UnixNano())
	fmt.Println(now.UTC().Unix())
	fmt.Println(now.UTC().UnixNano() / 1000000)
	fmt.Println(now.UTC().UnixNano())

	// 纳秒（19位）、微秒（16位）、毫秒（13位）、秒（10位）
	fmt.Println(now.UnixMicro())
	fmt.Println(now.UnixMilli())
	fmt.Println(now.UnixNano())
	
}
