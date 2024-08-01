package main

import (
	"fmt"
	"time"
)

func main() {

	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			fmt.Println("In goroutine for ever...")
			select {
			case <-done:
				fmt.Println("Stop ticker with done!")
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop() // 停止发送ticker
	time.Sleep(1000 * time.Millisecond)
	done <- true // 让goroutine从select中return 退出for循环
	// 如果不发送done ，select 将阻塞在done读取上
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("Ticker stopped")

}
