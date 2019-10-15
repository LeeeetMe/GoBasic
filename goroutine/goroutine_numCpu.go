package main

import (
	. "fmt"
	"runtime"
)

func getCpu(){
	// 获取CPU 个数，默认使用全部CPU
	cpu := runtime.NumCPU()
	Println("CPU Number is",cpu)
	// 控制最多使用多少cpu
	runtime.GOMAXPROCS(2)
}
