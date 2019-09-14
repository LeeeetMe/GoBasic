package base

import "fmt"

func Slice_test() {
	/*初始化区别
	slice：可以不指定切片长度，且长度可以更改
	array：必须指定数组长度，并不可更改
	*/
	var slice_one []int
	fmt.Printf("slice_one type is %T,value is ：%v\n", slice_one, slice_one)
	var arr1 [10]int
	fmt.Printf("arr1 type is %T,value is ：%v\n", arr1, arr1)

	// 切片取值
	var arr = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var slice_2 = arr[1:3]
	fmt.Printf("slice_2地址：%p,%v\n", &slice_2, slice_2)
	fmt.Printf("arr地址：%p,%v\n", &arr, arr)
	// slice没有自己的任何数据。它只是底层数组的一个表示。对slice所做的任何修改都将反映在底层数组中。
	for i, _ := range slice_2 {
		slice_2[i] += 1
	}
	fmt.Printf("slice_2地址：%p,%v\n", &slice_2, slice_2)
	fmt.Printf("arr地址：%p,%v\n", &arr, arr)
	// 两个切片公用一个底层数组时，修改切片就是修改底层数组
	var slice_3 = arr[4:8]
	for i, _ := range slice_3 {
		slice_3[i] += 1
	}
	fmt.Printf("slice_3地址：%p,%v\n", &slice_3, slice_3)
	fmt.Printf("arr地址：%p,%v\n", &arr, arr)

	// 切片长度：len(),切片容量cap()
	// make(类型, 长度,容量)
	var slice_4 = make([]int, 3, 5)
	fmt.Printf("arr_4切片长度：%d\n切片容量：%d\n切片类型：%p\n切片值默认为：%v\n", len(slice_4), cap(slice_4), slice_4, slice_4)
	var slice_5 = make([]int, 0, 5)
	fmt.Printf("slice_5切片长度：%d\n切片容量：%d\n切片类型：%p\n切片值默认为：%v\n", len(slice_5), cap(slice_5), slice_5, slice_5)
	// append 向slice里面追加一个或者多个元素，然后返回一个和slice一样类型的slice copy ，函数copy从源slice的src中复制元素到目标dst，并且返回复制的元素的个数
	slice_5 = append(slice_5, 10)
	fmt.Printf("slice_5切片长度：%d\n切片容量：%d\n切片类型：%p\n切片值默认为：%v\n", len(slice_5), cap(slice_5), slice_5, slice_5)
}
