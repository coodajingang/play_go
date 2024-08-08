package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	os.Setenv("foo", "bar") // 设置环境变量
	os.Setenv("x", "18")

	fmt.Println(os.Getenv("foo"))
	fmt.Println(os.Getenv("x"))
	fmt.Println(os.Getenv("x12341234"))
	fmt.Println(os.LookupEnv("x11234")) // 检查环境变量是否存在
	fmt.Println(os.LookupEnv("PATH"))
	fmt.Println(os.Getenv("PATH")) // 获取环境变量，可以是自己设置的 或 系统的环境变量

	for _, env := range os.Environ() { // 获取系统的环境变量
		s := strings.SplitN(env, "=", 2)
		fmt.Println(s[0], s[1])
	}
}
