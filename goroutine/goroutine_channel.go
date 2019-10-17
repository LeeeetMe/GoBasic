package main

import (
	. "fmt"
	"time"
)
/*
	初始化chan
*/
func InitChannel() chan int {
	// 初始化 channel是一个容器，定义后，要进行初始化操作
	//定义存储的数据是int类型的chan a
	var a chan int
	Printf("c = %#v\n", a)
	/* 初始化chan a,并指定a 的空间大小
	 如果没有指定空间大小：
		向chan内[存放/获取]数据时，会被阻塞，只有在向该chan[存放/获取]数据时，才会将数据[获取/存入]。
	 */
	a = make(chan int,10)
	//a = make(chan int)
	Printf("c = %#v\n", a)
	return a
}

/*
	存放数据，使用 channel <-
		channel <- data
*/
func produce(c chan int) {
	c <- 100
	Println("Produce 100 in c")
}
/*
	获取数据，使用 <- channel
		data := <- channel
*/
func consume(c chan int) {
	// 如果没有变量接收时，从channel c中取出一个数据并丢弃
	<-c
	//当队列中没有数据时，会阻塞，直到有数据存储进channel或主线程结束
	data := <-c
	Println("get Data from channel c:",data)
}

/*
	channel作为单向参数,
		<-在左边：只能获取 <-channel
		<-在右边：只能存储 channel<-
		函数内的<-要与参数的<-在同一侧
*/
func single(c chan<- int){
	//错误：传入的c 只能进行存储，不可以获取
	//data := <-c
}

func ChannelTest()  {
	c := InitChannel()
	/*
	没有声明channel 空间：
		存数据时，会阻塞，直到有人向该channel取数据时才会放行
	声明channel空间：
		存数据，取数据都不会阻塞
	*/
	go produce(c)
	go consume(c)
	time.Sleep(time.Second*5)
}


func ChannelIsClosed() {
	c := InitChannel()
	for i := 0;i<10;i++	 {
		c<-i
	}
	close(c)
	for {
		// 如果不进行channel判断，则会无线死循环
		//Println(<-c)

		//进行判断c是否关闭
		v,ok := <-c
		if !ok{
			Println("channel is closed")
			break
		}
		Println(v)
	}
}
