package base

import (
	. "fmt"
)

// 创建结构体
type Books struct {
	title   string
	author  string
	subject string
	book_id int
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

}
