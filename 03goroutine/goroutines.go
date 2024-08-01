package main

import (
	"fmt"
	"time"
)

func f(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s, " ", i)
	}
}

func main() {
	fmt.Println("Start ...")

	go f("goroutine")

	go func() {
		fmt.Println("goroutine2")
	}()

	time.Sleep(time.Millisecond)
	fmt.Println("End")
}
