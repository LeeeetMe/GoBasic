package main

import (
	. "fmt"
)

var (
	name string
	age  string
	sex  bool
)

func main() {
	Println("HelloWorld")
	Println("爱情是什么味道")
	name = "赵英涵"
	age = "28"
	sex = true

	//io_reader := bufio.NewReader(os.Stdin)
	//read_str, _ := io_reader.ReadString('\n')
	//Println(read_str)

	var x int
	var y string
	Scanln(&x, &y)
	Printf("x = %d 类型为：%T\n", x, x)
	//Printf("y = %s 类型为：%T\n", y, y)
	// 使用scanf时，需要注意的是 要按照字符串内的格式进行输入，如下：输入100,帅
	//Scanf("输入%d,%s", &x,&y)
	//Printf("x = %d ，y = %s", x, y)
	// if 语句使用
	//If_doc()
	/*  goto*/
	// basic_pro.Goto_obj()
	// 随机数
	// basic_pro.Random_obj()
	//
}