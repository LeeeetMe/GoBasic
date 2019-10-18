package main

import (
	. "fmt"
	"time"
	"sync"
)

func generateInt(i int,ch chan bool){
	time.Sleep(time.Second)
	Println("func generateInt",i)
	// 执行完一次，在ch中添加一个数值
	ch<-true
}

func waitGoroutine() {
	no := 3
	var c chan bool
	c = make(chan bool,no)
	for i := 0; i < no; i++ {
		go generateInt(i,c)
		<-c
	}
}

func waitAnything() {
	// 使用channel阻塞进程,当有capacity的channel中没有元素是调用<-ch时，会发生阻塞事件，
	// go一个时，添加一个元素
	// 在外部执行 no 次go 则取no次数值，则主线程就会等待goroutine线程结束后执行
	waitGoroutine()

	/*
	使用sync包，是主线程等待goroutine线程
	var sg sync.WaitGroup
	Add(num)  添加计数
	Done()  计数-1
	Wait()  等待计数归零后释放
	*/
	syncWait()
}

func syncGoroutine(i int,wg *sync.WaitGroup) {
	time.Sleep(time.Second)
	Println("syncGoroutine:",i)
	wg.Done()
}

func syncWait(){
	no :=3
	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		wg.Add(1)
		go syncGoroutine(i,&wg)
	}
	wg.Wait()
	Println("all goroutine is over")
}