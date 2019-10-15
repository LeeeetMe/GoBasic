package main

import (
	. "fmt"
	"time"
)

/*
Windows安装：
	1.执行 go get -u github.com/derekparker/delve/cmd/dlv
	2.上述命令执行后在gopath中的bin目录下会生成一个编译好的dlv.exe可执行程序
	3.将bin文件夹添加到环境变量后，终端执行dlv

基本使用：
	进入需要调试的包内
	dlv 命令 包名： dlv debug delveOnline
	执行后进入交互命令界面
	b(breakpoint缩写) 包名.函数名  执行断点调试
		c 执行
		next 下一步
		p(print) 变量名 ： p a -> 100
		r(restart) 重新执行
		s(step) 单步调试，想要进入嵌套函数内部时，可以使用

线上程序调试：
	dlv attach 程序PID
	dlv attach 11828

*/

func Add(a, b int) (sum int) {
	sum = a + b
	return
}

func RunAdd() {
	a, b := 100, 200
	c := Add(a, b)
	Println("a + b = ", c)
}

func LoopTime() {
	for {
		var i int
		var curTime time.Time
		time.Sleep(5 * time.Second)
		i++
		curTime = time.Now()
		Printf("run %d count, cur time is %v", i, curTime)
	}
}

func main() {
	LoopTime()
}
