package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://gobyexample-cn.github.io/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// readWithIoutil(resp)
	// readWithScanner(resp)
	reader := bufio.NewReader(resp.Body)
	buf := make([]byte, 1024)
	for {
		n, err := reader.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		os.Stdout.Write(buf[:n])
	}
}

func readWithScanner(resp *http.Response) {
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func readWithIoutil(resp *http.Response) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
