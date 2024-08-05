package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
func main() {

	bs, err := os.ReadFile("tmp2/dat2.txt")
	check(err)
	fmt.Println(string(bs))

	f, err := os.Open("tmp2/dat2.txt")
	check(err)
	defer f.Close()
	bs, err = ioutil.ReadAll(f)
	check(err)
	fmt.Println(string(bs))
	f.Seek(0, 0)
	b1 := make([]byte, 5)
	n, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n, b1[:n])

	n, err = f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n, b1[:n])
	f.Seek(20, 0)
	fmt.Println("After seeking...20,0")
	n, err = f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n, b1[:n])

	b2 := make([]byte, 20)
	// ReadAtLeast 是从当前文件位置读取至少指定的长度
	least, err := io.ReadAtLeast(f, b2, 20)
	check(err)
	fmt.Println("Read least:", least)
	fmt.Printf("%d bytes: %s\n", least, b2[:least])

	reader := bufio.NewReader(f)
	peek, err := reader.Peek(20)
	check(err)
	fmt.Println(string(peek))
	if err == io.EOF {
		fmt.Println("EOF Readed!")
	}

	f.Seek(0, 0)
	// 重新seek后，使用原有的reader 还是缓存部分以前读取的内容
	fmt.Println("Use bufio reader after seek 0,0")
	peek, err = reader.Peek(30)
	check(err)
	fmt.Println(string(peek))
	fmt.Println(len(peek))

	// 重新seek后，并且调用Reset后将从头开始读取
	f.Seek(0, 0)
	reader.Reset(f)
	fmt.Println("Use bufio reader reset ")
	peek, err = reader.Peek(30)
	check(err)
	fmt.Println(string(peek))
	fmt.Println(len(peek))
}
