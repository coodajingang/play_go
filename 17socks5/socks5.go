package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go func() {
			defer conn.Close()
			handle(conn)
		}()
	}
}

func handle(conn net.Conn) {
	for {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			msg := scanner.Text()
			fmt.Println("Received:", msg)
			n, err := conn.Write([]byte(msg + "\r\n"))
			if err != nil {
				panic(err)
			}
			fmt.Println("Send size: ", n)
		}
	}
}
