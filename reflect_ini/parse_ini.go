package main

import (
	. "fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
	// "errors"
)

type ConfigInfo struct {
	ServerConf Server `ini:"server"`
	MysqlConf  MySql  `ini:"mysql"`
}
type Server struct {
	Host string `ini:"host"`
	Port int    `ini:"port"`
}
type MySql struct {
	Username string `ini:"username"`
	Password string `ini:"password"`
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
}

func ReadFile(path string) []string {
	// 返回byte[]切片
	file, err := ioutil.ReadFile(path)
	if err != nil {
		Println("this has err：", err)
	}
	// 转换成列表返回
	strLst := strings.Split(string(file), "\n")
	return strLst
}

func ParseIni(str []string, result interface{}) (err error) {
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
	getType := reflect.TypeOf(result).Elem()
	if getValue.Kind() != reflect.Ptr {
		err = Errorf("传进来的result不是一个指针,%v", getValue.Kind())
	} else {
		Println("这是个指针")
	}
	if getValue.Elem().Kind() != reflect.Struct {
		err = Errorf("传进来的result不是一个结构体,%v", getValue.Elem())
	} else {
		Println("这是个结构体")
	}
	// 声明节点变量
	var lastFieldName string
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
			lastSectionName := line[1 : len(line)-1]
			Println("节点名称为", lastSectionName)
			lastFieldName, err = getFieldName(lastSectionName, getType)
			if err != nil {
				Println("这是个错误啊")
				return err
			}
			Println("lastFieldName is ", lastFieldName)
			continue
		} else if strings.Contains(line, "=") && strings.Count(line, "=") == 1 {
			Println("嗯，这不是节点这是一个节点内的元素", line)
		} else {
			Printf("小朋友不要乱按键盘，输入有问题，值：%s行号：%d\n", line, index+1)
		}
		err = parseItem(lastFieldName, line, result)
		if err != nil {
			return
		}
	}
	return
}

func getFieldName(lastSectionName string, typeInfo reflect.Type) (fieldName string, err error) {
	for index := 0; index < typeInfo.NumField(); index++ {
		fieldType := typeInfo.Field(index)
		tagValue := fieldType.Tag.Get("ini")
		if tagValue == lastSectionName {
			fieldName = fieldType.Name
			Println("返回的FiledName is", fieldName)
			break
		}
	}
	if fieldName == "" {
		err = Errorf("no field name is %s", lastSectionName)
	}
	return
}

func parseItem(lastFieldName string, line string, result interface{}) (err error) {
	Println(lastFieldName, line)
	index := strings.Index(line, "=")
	if index == -1 {
		err = Errorf("syntax error in line :%s", line)
		return
	}
	key := strings.TrimSpace(line[0:index])
	val := strings.TrimSpace(line[index+1:])
	if len(key) == 0 {
		err = Errorf("sytax error, line:%s", line)
		return
	}
	Println(key, val)
	resultValue := reflect.ValueOf(result)
	sectionValue := resultValue.Elem().FieldByName(lastFieldName)
	sectionType := sectionValue.Type()
	Println("sectionValue is", sectionValue)

	if sectionValue.Kind() != reflect.Struct {
		err = Errorf("field:%s must be struct", lastFieldName)
		return
	}
	var itemName string
	itemName, err = getFieldName(key, sectionType)
	if err != nil {
		Println(err)
		return
	}
	itemValue := sectionValue.FieldByName(itemName)
	switch itemValue.Kind() {
	case reflect.String:
		itemValue.SetString(val)
	case reflect.Int8, reflect.Int16, reflect.Int, reflect.Int32, reflect.Int64:
		intVal, errRet := strconv.ParseInt(val,10,64)
		if errRet != nil{
			err = errRet
			return
		}
		itemValue.SetInt(intVal)
	}
	Println("itemName is", itemName)
	Println("itemValue is", itemValue)
	return
}
