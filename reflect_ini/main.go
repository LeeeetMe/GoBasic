package main

import (
	. "fmt"
)

func main() {
	var configIni ConfigInfo
	str := ReadFile("conf.ini")
	Println(str)
	_ = ParseIni(str, &configIni)
	println("===========")
	Printf("%#v",configIni)
}
