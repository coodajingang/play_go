package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
func main() {

	err := os.Mkdir("tmp2", os.ModeDir+os.ModePerm)
	check(err)

	err = os.Mkdir("tmp3", 0644)
	check(err)

	// 读取目录下所有的条目
	dir, err := os.ReadDir("14file")
	check(err)
	for _, d := range dir {
		fmt.Println(d.Name())
		fmt.Println(d.IsDir())
		fmt.Println(d.Type())
		fmt.Println(d.Info())
	}

}
