package base

import (
	. "fmt"
	"strconv"
	"strings"
	s "strings"
)

func String_obj() {
	/*
		字符串是：
			多个字符组成的序列,
			字符串不可更改
		字符：
			概念：对应编码表中的编码
				A ---> 65
				B ---> 66
				a ---> 97
				对字符A的操作，就是操作她的编码65
			类型：
				字符的类型就是字节--> byte --> utf-8
				中文字符在utf-8中，占3个字节
				字符可以对byte进行转换

	*/
	s1 := "爱吃馒头喝牛奶"
	Printf("%s %T \n", s1, s1)
	Printf("%c %T \n", s1[0], s1[0])

	//1.一个中文字符占3个字节,
	// for i := 0; i < len(s1); i++ {
	// 	Printf("%c %T \n", s1[i], s1[i])
	// }
	//1.字符可通过 %c 格式化输出
	s2 := 'A'
	Println(s2)
	Printf("%c\n", s2)
	// 2.字符串与byte[]转换
	s3 := []byte{65, 66, 67, 68}
	s4 := string(s3)
	Println(s4)

	s5 := "ABCD"
	s6 := []byte(s5)
	Println(s6)
	// 3.字符串不可以修改
	// 修改s5
	// s5[3] = 'd'  //cannot assign to s5[3]

	/*
		string 常用函数
		官方文档：https://golang.google.cn/pkg/strings/
	*/
	s7 := "hello world"
	s8 := "held"
	s9 := "l"
	//  s7 是否包含 s8
	is_contains := s.Contains(s7, s8)
	Println(is_contains)
	// s7 和 s8 是否有相同的字符
	is_contains = s.ContainsAny(s7, s8)
	Println(is_contains)
	// 统计s9 在s7中出现的次数并返回
	num := s.Count(s7, s9)
	Println(num)
	// 判断字符串以...开头 ...结尾
	s10 := "golang基础.txt"
	s11 := "golang"
	s12 := "txt"
	if s.HasPrefix(s10, s11) {
		Println("这个是golang相关文件")
	}
	if s.HasSuffix(s10, s12) {
		Println("这是txt文件")
	}
	// Index
	s13 := "checken"
	// index 中的substirng需要完全匹配才会返回index，否则返回-1
	Println(s.Index(s13, "ken"))
	Println(s.Index(s13, "dmr"))
	// 如果substring中在字符串中出现多次，只会返回第一次出现的index
	Println(s.Index(s13, "e"))

	// substring与s13相同的字符，在s13中第一次出现的index
	Println(s.IndexAny(s13, "dfdfe"))
	// substring 最后一次出现的索引位置
	Println(s.LastIndex(s13, "e"))

	// 字符串拼接join
	s14 := []string{"hello", "my", "world"}
	s15 := s.Join(s14, "-")
	Println(s15)
	// 重复拼接
	s16 := s.Repeat(s14[1], 3)
	Println(s16)
	// 字符串切割
	s17 := s.Split(s15, "-")
	Printf("%T,%v\n", s17, s17)
	// 替换
	Println(strings.Replace("oink oink oink", "k", "ky", 2))      //2 代表只替换前两个
	Println(strings.Replace("oink oink oink", "oink", "moo", -1)) // -1代表替换所有
	// map 每一个元素都执行绑定的函数

	s18 := "Alex like DDD"
	s19 := s.Map(UpAndLower, s18)
	Println(s19)
	Printf("%d\n%d\n", 'D', 'd')

	// string 中的 strconv 函数，用于string与各个类型之间的转换
	// boolean 类型转换
	// format... 格式化为字符串
	val := strconv.FormatBool(true)
	Printf("%s  %T \n", val, val)
	// parse...  从字符串解析成各种类型
	flag, _ := strconv.ParseBool(val)
	Printf("%t  %T \n", flag, flag)
	// 数值之间的转换
	// string -> int
	s20 := "100"
	//s20为源字符串，10为s20表示的几进制，64代表系统32/64位
	int_val, _ := strconv.ParseInt(s20, 10, 64)
	Printf("%d  %T \n", int_val, int_val)
	s21 := strconv.FormatInt(int_val, 10)
	Printf("%s  %T \n", s21, s21)
	// Atoi、Itoa ,整数之间的转换
	sValue := "-100"
	iValue, _ := strconv.Atoi(sValue)
	Printf("%d  %T \n", iValue, iValue)
	s22 := strconv.Itoa(iValue)
	Printf("%s  %T \n", s22, s22)

}

func UpAndLower(r rune) rune {
	switch {
	case r >= 'A' && r <= 'Z':
		return r + 32
	case r >= 'a' && r <= 'z':
		return r - 32
	}
	return r
}
