package base

import (
	. "fmt"
)

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

type Person struct {
	int
	string
}

func Struct_test() {
	/*
		结构体：
			由一系列具有相同各类型或不同类型构成的数据集合 (class)
		拷贝：
			值类型数据，需要使用指针进行拷贝
	*/
	// 1
	var b1 Books
	b1.author = "鲁迅"
	b1.title = "狂人日记"
	b1.subject = "123"
	b1.book_id = 1
	Println(b1)
	// 2
	var b2 Books = Books{}
	b2.author = "苏轼"
	b2.title = "江城子"
	b2.subject = "666"
	b2.book_id = 2
	Println(b2)
	// 3
	b3 := Books{
		author:  "鲁迅",
		title:   "朝花夕拾",
		subject: "222",
		book_id: 3,
	}
	Println(b3)
	// 4
	b4 := Books{"如梦令", "李清照", "抒情", 4}
	Println(b4)
	/*
		匿名字段：
			1.当字段没有名字时，直接使用字段类型表示该字段的值
			2.在同一结构体中匿名字段不可以重复
	*/
	var p Person = Person{
		11,
		"单身狗",
	}
	Println(p.int, p.string)

}

func Inherit_struct() {
	/*
		继承：结构体的嵌套
			1.在结构体中通过匿名字段的方式将其他结构体嵌入进来，可以直接通过.的方式访问嵌套的字段
	*/
	// 1
	type Food struct {
		category string
		weight   int
	}

	type Chips struct {
		price int
		f     Food
		num   int
	}

	type Noodles struct {
		Food
		price    int
		docoment string
	}

	var n Noodles
	n.category = "主食"
	n.weight = 200
	n.price = 5
	n.docoment = "你看这面，又多又好吃"
	Println(n)
	// 2
	var c Chips = Chips{
		f: Food{
			category: "零食",
			weight:   180,
		},
		price: 10,
		num:   2,
	}
	Println(c, c.f.category, c.f.weight, c.price, c.num)
}

type People struct {
	name string
	age  int
}

type Children struct {
	People
	interest string
	address  string
}

func (p People) run() {
	Printf("我是父类的Run，%s 正在奔跑，他已经%d岁了\n", p.name, p.age)
}
func (c Children) run() {
	Printf("你好，我是%s今年%d岁了，我家在%s\n", c.name, c.age, c.address)
}
func (c Children) sayHello() {
	Printf("我是子类的Run你好，我是%s今年%d岁了，我家在%s\n", c.name, c.age, c.address)
}

func Struct_function() {
	/*
		struct之间的嵌套子类，父类：
			父类不可以调用子类的函数
			子类可以调用父类的函数，子类也可以重写父类的方法，调用时使用就近原则
	*/
	var c1 Children = Children{
		People: People{
			name: "alex",
			age:  100,
		},
		address: "东北",
	}
	c1.sayHello()
	c1.run()
}
