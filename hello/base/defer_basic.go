package base

import (
	. "fmt"
)

/*
defer：延迟，推迟
1.方法：一个函数被延迟【执行】
2.用法：
	A：对象.close()，临时文件的删除。
		文件.open()
		defer 文件.close()
		可用于场景切换 释放资源
	B：异常处理：panic()、recover()
		panic()：程序执行被中断
		recover()：恢复程序执行，recover()必须在defer中执行
3. 多个defer时，遵循栈规则，先进后出
4. defer 传参时需要注意：defer函数调用时，已经传递参数了,只是没有【执行】
5.注意事项：
	A：所有函数执行完成后执行defer
	B：同一代码块中有return语句时，所有defer都执行完，才会执行return语句
*/

func manyDefer() string {
	Println("1")
	defer Println("defer 1")
	Println("2")
	Println("3")
	defer Println("defer 2")
	return "alex"
}

func DeferBasic() {
	Println(manyDefer())
}
