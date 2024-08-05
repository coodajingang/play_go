package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	fmt.Println(e)
}
func main() {
	// 写文件，不存在创建之， 然后truncate 写入文件 （覆盖写）
	err := os.WriteFile("tmp2/dat1.txt", []byte("hello wor"), 0644)
	check(err)

	// 创建文件，不存在则使用0666创建之， 存在则truncate之
	f, err := os.Create("tmp2/dat2.txt")
	check(err)
	defer f.Close()
	f.WriteString("hello wor \n")
	f.WriteString("this is line 2 \n")
	fmt.Fprintln(f, "This is line 3 \n")

	f.Sync()

	writer := bufio.NewWriter(f)
	writer.WriteString("this is line 4 \n")
	writer.Flush()

	stat, err := f.Stat()
	check(err)
	fmt.Println(stat.Size())
	fmt.Println(stat.ModTime())
	fmt.Println(stat.Mode())
	fmt.Println(stat.Mode().IsRegular())
}
