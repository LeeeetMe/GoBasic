package base

import (
	. "fmt"
)

/*
	指针变量：存放地址的变量(一个变量内存储的值，是另一个变量的地址，这样的变量叫做指针变量)
	& ：取地址符
	* ：取地址内的值
*/

func Pointer_value() {
	// & 取地址
	var i int = 1
	Println("变量i的地址是", &i)

	// 指针变量的声明，赋值
	// 创建一个存储类型为int
	var a *int = &i
	Printf("a的指针值：%v，类型：%T ,指针变量a地址：%p\n", a, a, &a)

	// 相同类型的指针变量才可以赋值
	var b *int = a
	Printf("b的指针值：%v，类型：%T ,指针变量b地址：%p\n", b, b, &b)

	// * 取地址内的值
	Printf("指针变量a存储的地址中的值为%d\n", *a)
	Printf("指针变量b存储的地址中的值为%d\n", *b)
	// ** 存储指针变量地址的指针
	// c := &a
	var c **int = &a
	Printf("c的指针值：%v，类型：%T ,指针变量c地址：%p\n", c, c, &c)
	// 取值
	Printf("指针变量c存储的地址中的地址中的值为%d\n", **c)

}

func Pointer_array() {
	/*
		数组指针：存储数组地址的变量
	*/
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	// 创建一个指针，存储该数组的地址
	var arr_pointer *[5]int
	arr_pointer = &arr
	Printf("arr_pointer的指针值：%v，类型：%T ,指针变量地址：%p\n", arr_pointer, arr_pointer, arr_pointer)
	// 通过* 对数组内的值进行更改,通过* 取到地址内的数组后进行更改
	(*arr_pointer)[1] = 100
	Printf("arr_pointer的指针值：%v，类型：%T ,指针变量地址：%p\n", arr_pointer, arr_pointer, arr_pointer)
	// 简化写法
	arr_pointer[1] = 1000
	Printf("arr_pointer的指针值：%v，类型：%T ,指针变量地址：%p\n", arr_pointer, arr_pointer, arr_pointer)

	/*
		指针数组：数组内的值为指针
	*/
	a, b, c, d := 1, 2, 3, 4
	var arr_int [4]int = [4]int{a, b, c, d}
	var arr_p [4]*int = [4]*int{&a, &b, &c, &d}
	Printf("arr_int的指针值：%v，类型：%T ,指针变量地址：%p\n", arr_int, arr_int, &arr_int)
	Printf("arr_p的指针值：%v，类型：%T ,指针变量b地址：%p\n", arr_p, arr_p, &arr_p)

	// 更改arr_int中的变量的值,替换了a
	arr_int[0] = 100
	Printf("arr_int的指针值：%v，类型：%T ,指针变量地址：%p\n", arr_int, arr_int, &arr_int)
	Println("a的值:", a)
	// 更改var_p地址内的值 a 中的值也会被更改
	*arr_p[0] = 200
	Println("更改后a的值:", a)
}

func arr_func() [4]int {
	arr := [4]int{1, 2, 3, 4}
	return arr
}
func point_func() *[4]int {
	arr := [4]int{1, 2, 3, 4}
	return &arr
}

func Point_func() {
	/*
		指针函数：返回值是类型是指针的函数
	*/

	a := arr_func()
	p := point_func()
	Printf("a的指针值：%v，类型：%T ,指针变量地址：%p\n", a, a, &a)
	Printf("p的指针值：%v，类型：%T ,指针变量地址：%p\n", p, p, &p)
}
