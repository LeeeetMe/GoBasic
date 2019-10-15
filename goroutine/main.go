package main

import (
	. "fmt"
	"time"
)

func hello(i int) {
	Println("Hello goroutine", i)
}
func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		Printf("%d ", i)
	}
}

func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		Printf("%c ", i)
	}
}

func goroutine1()  {
	for index := 0; index < 10; index++ {
		go hello(index)
	}
}

func goroutine2()  {
	go numbers()
	go alphabets()
	time.Sleep(3000 * time.Millisecond)
	Println("main terminated")
}

func main() {
	//goroutine2()
	//goroutine1()
	//time.Sleep(13000 * time.Millisecond)
	//getCpu()
	//初始化channel后添加数据，再取数据
	//ChannelTest()
	//判断channel是否关闭
	ChannelIsClosed()
	Println("main thread terminate")
}
