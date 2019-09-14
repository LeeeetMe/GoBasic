package base

import (
	. "fmt"
)

/*
	函数参数类型：
	* 基本数据类型：
		int、string、bool、float
	* 复合数据类型：
		array，slice，map，function，pointer，struct，interface...
	函数的可变参数
	语法： 参数名 ...类型 --> 类似python中的*args
	* 可输入参数为多个参数时，最后输入可变参数
	* 可变参数只能有一个
	 传入切片时，需要在切片末尾使用 ... 来解析切片

*/

func VariableParameter() {
	Println("传入多个值")
	a_funcs(1, "我是B", 3, 4, 5, 6, 7)
	Println("传入切片")
	a_funcs(1, "我是B,后边那个要用三个点...解析切片了", []int{3, 4, 5, 6, 7}...)
}

func a_funcs(a int, b string, args ...int) {
	Printf("a:%d b:%s args：%v\n", a, b, args)
}

/*
	函数的值传递与引用传递
	值传递：struct，array
	引用传递：slice
*/
func ValAndRefParameter() {
	var a = [4]int{1, 2, 3, 4}
	var b = []int{1, 2, 3, 4, 5}
	Println(a)
	valParameter(a)
	Println(a)
	refParameter(b)
	Println(b)
}

// 值传递
func valParameter(nums [4]int) {
	nums[0] = 100
}

// 引用传递
func refParameter(nums []int) {
	nums[0] = 100
}

/*
	闭包：
		1.一个内层函数中调用外层函数创建的变量，且内层函数 return 这个变量
		2.外层函数 return 内层函数
*/
//闭包函数
func Increment() func() int {
	i := 1
	res1 := func() int {
		i++
		return i
	}
	return res1
}

func IncrementTest() {
	res1 := Increment()
	v1 := res1()
	v2 := res1()
	Println("v1 = ", v1)
	Println("v2 = ", v2)
}
