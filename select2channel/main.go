package main

import (
	. "fmt"
	"time"
)

func server1(ch chan string) {
	time.Sleep(time.Second * 2)
	ch <- "server1"
}

func server2(ch chan string) {
	time.Sleep(time.Second * 4)
	ch <- "server2"
}

func normalChannel() {
	/*不使用select的情况下：
	server1 需要2秒将数据放入ch1
	server2 需要4秒将数据放入ch2
	先执行<-ch2 需要四秒，<-ch1 原本只需要2秒，却要等待<-ch2 ,select 就是解决这个问题的
	*/
	ch1 := make(chan string, 10)
	ch2 := make(chan string, 10)
	go server1(ch1)
	go server2(ch2)

	Println(<-ch2)
	Println(<-ch1)
}

func selectChannel() {
	ch1 := make(chan string, 10)
	ch2 := make(chan string, 10)
	go server1(ch1)
	go server2(ch2)
	select {
	case s1 := <-ch1:
		Println(s1)
	case s2 := <-ch2:
		Println(s2)
	default:
		Println("default：所有的channel都为空")
	}
}
func putIntoChannel(ch chan string)  {
	for  {
		ch <- "hello"
	}
}
func getChannel(ch chan string) {
	time.Sleep(time.Second)
	for value := range ch {
		Println("get value from ch:",value)
	}
}
func selectFullChannel()  {
	/*
		检测channel是否已满
	*/
	ch := make(chan string,10)
	go putIntoChannel(ch)
	go getChannel(ch)
	for{
		select {
		case ch <- "good":
			Println("ch input good")
		default:
			Println("ch已经满了")
		}
	}
}

func main() {
	//normalChannel()
	//selectChannel()
	selectFullChannel()
}
