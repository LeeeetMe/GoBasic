package base

import (
	. "fmt"
)

func Array_test() {
	/* 数组初始化的几种方式 */
	// 1
	var arr [5]int //等价于  arr:= [5]int[0,0,0,0,0]
	Println(len(arr))
	Println(arr)

	// 2
	// 使用...时，初始化的长度为
	var arr1 = [...]string{"alex", "alice", "lilei", "hanmeimei"}
	Println(len(arr1))
	// 3
	// 如果数组内的数量小于规定的数量，则剩余的初始化默认为0
	arr2 := [10]int{1, 2, 3, 4, 5}
	Println(arr2)
	Println(len(arr2))



	/* 遍历数组 */
	// 1
	for i := 0; i < len(arr2); i++ {
		Printf("index：%d,value:%d\n", i, arr2[i])
	}
	// 2
	for _, val := range arr1 {
		Printf("name is %s\n", val)
	}


	/* 多维数组 */
	// [有几个数组][每个数组内有几个数]
	a := [3][4]int{
		{0, 1, 2, 3},   /*  第一行索引为 0 */
		{4, 5, 6, 7},   /*  第二行索引为 1 */
		{8, 9, 10, 11}, /*  第三行索引为 2,最后一个元素后如果没有}，需要要有， */
	}
	Println("第1行的第2个数", a[0][1])


	/* 数组是值类型 */
	arr3 := [...]string{"a", "b", "c", "d"}
	arr4 := arr3 // 将arr3拷贝后赋值给arr4
	arr4[0] = "alex"
	Println("arr3 is ", arr3)
	Println("arr4 is ", arr4)
	// 因此[5]int和[25]int是不同的类型。因此，数组不能被调整大小。不要担心这个限制，因为切片的存在是为了解决这个问题。
	var x [3]int
	// var b = [5]int{1, 2, 3, 4, 5}
	// x = b  cannot use b (type [5]int) as type [3]int in assignment
	Println(x)
}
