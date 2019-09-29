package main

import (
	. "fmt"
	"io/ioutil"
	"reflect"
	"strings"
)

type ConfigInfo struct {
	ServerConf Server `ini:server`
	MysqlConf  MySql  `ini:mysql`
}
type Server struct {
	host string
	port int
}
type MySql struct {
	user     string
	password string
	host     string
	port     int
}

func ReadFile(path string) []string {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		Println("this has err：", err)
	}
	strLst := strings.Split(string(file), "\n")
	return strLst
}

func ParseIni(str []string, result interface{}) {
	/*
		1.判断result是否为结构体指针
		2.遍历字符串数组
			1.去除所有空格，如果长度=0则从数组中删除
			2.判断元素首字符是否为[，且结尾字符为]，符合条件则这个就是节点
				1.判断是否包含“=”，不包含则报错
					包含则进行切割
						....
	*/
	getValue := reflect.ValueOf(result)
	if getValue.Kind() != reflect.Ptr {
		Errorf("传进来的result不是一个指针,%v", getValue.Kind())
	} else {
		Println("这是个指针")
	}
	if getValue.Elem().Kind() != reflect.Struct {
		Errorf("传进来的result不是一个结构体,%v", getValue.Elem())
	} else {
		Println("这是个结构体")
	}
	// 声明节点变量
	var lastSection string
	for index, val := range str {
		// 去除前后空格
		line := strings.TrimSpace(val)
		if len(line) == 0 {
			continue
		}
		if strings.HasPrefix(line, "#") || strings.HasPrefix(line, ";") {
			Println("这就是注释", line)
			continue
		}
		// 判断是否为开头[，]结尾 且不为空
		if len(line) > 2 && strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			Println("这是一个节点")
			lastSection := line[1 : len(line)-1]
			Println("节点名称为", lastSection)
		} else if strings.Contains(line, "=") && strings.Count(line, "=") == 1 {
			Println("嗯，这不是节点这是一个节点内的元素", line)
		} else {
			Printf("小朋友不要乱按键盘，输入有问题，值：%s行号：%d\n", line, index+1)
		}
		Println(lastSection)
	}
}
