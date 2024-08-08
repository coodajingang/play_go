package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {

	wFlag := flag.String("w", "default w value", "this is a w flag")
	var w2Flag string
	flag.StringVar(&w2Flag, "w2", "default w2 value", "this is a w2 flag")
	iFlag := flag.Int("i", 999, "this is a i flag")

	duration := flag.Duration("time", 10*time.Second, "this is a time flag")

	b := flag.Bool("bool", false, "this is a bool flag")

	/**
	-flag
	-flag=x
	-flag x  // 仅限非布尔标志
	使用- 或 --
	*/
	flag.Parse()
	fmt.Println(*wFlag)
	fmt.Println(*iFlag)
	fmt.Println(*duration)
	//
	fmt.Println(*b) // ./flags -bool=f  bool 类型的值为 1, 0, t, f, T, F, true, false, TRUE, FALSE, True, False

	fmt.Println(w2Flag)

	// Args 是其他的位置参数，需要放在所有的标志参数后面
	//返回命令行参数后的其他参数
	fmt.Println(flag.Args())
	//返回命令行参数后的其他参数个数
	fmt.Println(flag.NArg())
	//返回使用的命令行参数个数
	fmt.Println(flag.NFlag())
}
