package base

import (
	. "fmt"
)

type usb interface {
	start()
	end()
}

type mouse struct {
	name string
}

type keyboard struct {
	name string
}

type cat struct {
	name string
}

func (m mouse) start() {
	Println("我的名字是", m.name)
}
func (m keyboard) start() {
	Println("我的名字是", m.name)
}
func (m mouse) end() {
	Println("我的功能是AAA")
}
func (m keyboard) end() {
	Println("我的功能是QWERDF")
}

func StartToEnd(u usb) {
	u.start()
	u.end()
}

func Interface_obj() {
	/*
		接口：
			type 接口名字 interface{
				fun_name(parameter)
			}
		- 1.当某个struct 为这个接口的【所有方法】提供了方法实现，它就是实现接口的类
		- 2.当需要接口类型的对象时，可以使用任意实现接口类进行代理
	*/
	// 1
	Println("this is Interface")
	var m mouse = mouse{name: "蝰蛇"}
	var k keyboard = keyboard{name: "FILCO"}
	m.start()
	m.end()
	k.start()
	k.end()

	// 2
	// 实现接口类赋值给接口变量后，接口变量u_m 只能调用当前接口的方法，不能调用实现接口类
	var u_m usb = m
	u_m.start()
	u_m.end()
	// 如果一个函数需要传入接口类型对象作为参数，那么可以传入任意实现这个接口的对象实例作为参数
	StartToEnd(m)
	StartToEnd(k)
}

/*
	空接口：因为没有任何方法的接口，表示所有的类型都实现了空接口，因此空接口可以存储任意类型的数值
*/
func EmptyInterface() {
	type empty interface{}
	var a empty = "alex"
	var b empty = 10
	var c empty = [4]int{1, 2, 3, 4}
	var d empty = []int{1, 2, 3, 4}
	var e empty = cat{name: "花猫"}
	Println(a)
	Println(b)
	Println(c)
	Println(d)
	Println(e)

	// map
	var m map[string]interface{} = make(map[string]interface{})
	m["name"] = a
	m["age"] = b
	Println(m)

	// slice
	var s []empty = make([]empty, 0, 10)
	s = append(s, a)
	s = append(s, b)
	s = append(s, c)
	s = append(s, d)
	Println(s)
}

func getType(u usb) {
	switch x := u.(type) {
	case mouse:
		Println("我是鼠标", x.name)
	case keyboard:
		Println("我是键盘", x.name)
	default:
		Println("什么类型都算不上")
	}
}
func InterfaceType() {
	var m mouse = mouse{name: "我是苹果鼠标"}
	var k keyboard = keyboard{name: "我是苹果键盘"}
	getType(m)
	getType(k)
}
