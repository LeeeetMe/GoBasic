package base

import (
	. "fmt"
)

/*
	接口嵌套
	接口对象不能调用
*/
type A interface {
	test1()
}

type B interface {
	test2()
}

type C interface {
	A
	B
	test3()
}

type cat struct {
	name string
}

func (c cat) test1() {
	Println("test1 name is ", c.name)
}

func (c cat) test2() {
	Println("test2 name is ", c.name)
}

func (c cat) test3() {
	Println("test3 name is ", c.name)
}

func NestingInterface() {
	var c cat = cat{name: "Tom"}
	c.test1()
	c.test2()
	c.test3()

	var c1 C = c
	c1.test1()
	c1.test2()
	c1.test3()
	// 如果声明 A 接口变量，将接口实现类cat的实例c赋值给a1，则a1只能调用c中的test1方法，因为A接口中只有test1方法
	var a1 A = c
	a1.test1()
	var b1 B = c
	b1.test2()
	// 如果通过实现接口类型为A的变量赋值给 实现接口类型为C 的变量是不行的，好比把接口A类型看做是接口C类型，是不可以的，因为C>A
	// cannot use a1 (type A) as type C in assignment:A does not implement C (missing test2 method)
	// var c2 C = a1

}
