package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	dial, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	defer dial.Close()
	reader := bufio.NewReader(os.Stdin)
	go func() {
		for {
			scanner := bufio.NewScanner(dial)
			for scanner.Scan() {
				fmt.Println(scanner.Text())
			}
		}
	}()
	for {
		readString, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		n, err := fmt.Fprint(dial, readString)
		if err != nil {
			panic(err)
		}
		if n != len(readString) {
			fmt.Println("n != len(readString)")
		}
		fmt.Println("Write ", n)

	}
}
