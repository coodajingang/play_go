package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.NewTimer(2 * time.Second)
	<-t.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		fmt.Println("goroutine blocked with timer2")
		<-timer2.C
		fmt.Println("timer 2 fired")
	}()

	// 下面这段sleep 让timer2在stop前有机会触发
	fmt.Println("Sleep main with 2s")
	time.Sleep(2 * time.Second)

	// 若stop成功，timer将没有机会fire，goroutine将一直阻塞等待timer信号
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stoped!", stop2)
	} else {
		// 若stop失败，说明timer已经触发过了
		fmt.Println("Timer 2 already fired!")
	}
	time.Sleep(2 * time.Second)

	fmt.Println("Main 将阻塞1s 等待after结束继续执行。", time.Now())
	a := time.After(1 * time.Second)
	<-a
	fmt.Println("Main 阻塞1s after结束继续执行。", time.Now())

	fmt.Println("Main 将阻塞1s 等待after结束执行函数。", time.Now())

	// AfterFunc 将在等待结束后自己当前的goroutine里执行函数
	// 其返回的timer用来stop自己 ， 其返回的C是nil，在其上读会造成deadlock
	timer3 := time.AfterFunc(1*time.Second, func() {
		fmt.Println("AfterFunc 后执行的函数", time.Now())
	})
	fmt.Printf("%#v\n", timer3)
	//<-timer3.C // timer3.C == nil  read it will deadlock
	fmt.Println("Main 阻塞1s after结束函数继续执行。", time.Now())

	time.Sleep(2 * time.Second)
	fmt.Println("end main")

}
