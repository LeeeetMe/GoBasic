package main

import (
	. "fmt"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		Println("err is", err)
		return
	}
	conn, err := listen.Accept()
	if err != nil {
		Println("err is", err)
		return
	}
	var buf [128]byte
	for {
		n, err := conn.Read(buf[:])
		if err != nil {
			Println("err is", err)
			break
		}
		Println(string(buf[:n]))
	}
}
