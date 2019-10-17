package main

import (
	. "fmt"
)

func channelLength() {
	ch := make(chan int, 10)
	for i:=0;i<5 	;i++  {
		ch<-i
	}
	Println("channel length is",len(ch))
	ch<-100
	Println("channel length is",len(ch))
	Println("channel cap is",cap(ch))
}
func main() {
	//channel
	//channelLength()
	//主线程等待goroutine
	waitAnything()
}