package main

import (
	. "fmt"
)

func main() {
	var configIni ConfigInfo
	str := ReadFile("conf.ini")
	Println(str)
	ParseIni(str, &configIni)
}
