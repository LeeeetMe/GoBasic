package main

import (
	"bufio"
	. "fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		Println("err is ", err)
	}
	//控制台输入
	render := bufio.NewReader(os.Stdin)
	for {
		data, err := render.ReadString('\n')
		if err != nil {
			Println("err is ", err)
			return
		}
		data = strings.TrimSpace(data)
		_, err = conn.Write([]byte(data))
		if err != nil {
			Println("err is ", err)
			break
		}
	}

}
