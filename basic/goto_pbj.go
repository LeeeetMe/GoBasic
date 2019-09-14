package basic_pro

import (
	. "fmt"
)

func Goto_obj() {
	// goto语句通过设置标签，通过关键字goto 进行跳转
	for i := 0; i < 10; i++ {
		Println(i)
		if i == 5 {
			goto breakHere
		}
	}

breakHere:
	Println("跳转到我这里了")
}
