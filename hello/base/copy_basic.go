package basic_obj

import (
	. "fmt"
)

func Copy_test() {
	/*
		copy  只会将slice第一层的值进行拷贝，如嵌套slcice内的值，则都会被改变。
	*/
	arr_1 := [10]int{1, 2, 3, 4, 5, 6, 7, 8}
	slice_0 := make([][]int, 10)
	slice_1 := arr_1[:5]
	slice_2 := arr_1[5:]
	slice_0 = append(slice_0, slice_1, slice_2)
	Println("slice_0 is", slice_0)
	Println("slice_1 is", slice_1)
	Println("slice_2 is", slice_2)
	slice_3 := make([][]int, 20)
	copy(slice_3, slice_0)
	Println("s3 is", slice_3)
	for i, _ := range slice_1 {
		slice_1[i] += 1
	}
	Println("s1 is", slice_1)
	Println("s3 is", slice_3)

	// 如果一个copy(slice(小)，slice(大)) 会将slice共同位置的值进行拷贝
	slice_4 := []int{1, 2, 3, 4}
	slice_5 := make([]int, 2)
	copy(slice_5, slice_4)
	Println("slice_5 is", slice_5)
	// copy时，也可以直接进行切片拷贝(对切片进行扩容)
	copy(slice_5[1:], slice_4[3:])
	Println("slice_5 is", slice_5)
	// slice copy
	/*
		Slice在多次append元素时，若满足扩容策略，这时候内部就会重新申请一块内存空间，
		将原本的元素拷贝一份到新的内存空间上。此时其与原本的数组就没有任何关联关系了，
		再进行修改值也不会变动到原始数组。
	*/
	slice_6 := []int{1, 2, 3, 4}
	slice_7 := slice_6
	slice_7[0] = 9
	Println("slice_6 is", slice_6)
	Println("slice_7 is", slice_7)
	slice_8 := make([]int, 4)
	copy(slice_8, slice_6)
	slice_8[0] = 1
	Println("slice_6 is", slice_6)
	Println("slice_8 is", slice_8)
}
