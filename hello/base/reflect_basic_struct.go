package base

import (
	. "fmt"
	"reflect"
)

func get_value_type(f interface{}) {
	t := reflect.TypeOf(f)
	v := reflect.ValueOf(f)
	Println("类型为", t)
	Println("类值型为", v)

	switch t.Kind() {
	case reflect.Float32:
		Println("this is Float32")
	case reflect.Int:
		Println("this is Int")
	case reflect.Ptr:
		Println("this is Pointer")
		// 如果传入的是指针，就需要通过v.Elem()来设置值
		// Elem() 就相当于对v这个指针变量进行取值 *v
		v.Elem().SetFloat(6.66)
	default:
		Println("what is this ？")
	}
}

func Reflect_test() {
	// var f float32 = 3.12
	// get_value_type(&f)
	// Println(f)
	StructReflect()
}

/*
	结构体使用reflect
*/

type Phone struct {
	Color string
	Brand string
	Price int
}

func StructReflect() {
	var p Phone
	v := reflect.ValueOf(p)
	t := v.Type()
	k := t.Kind()
	switch k {
	case reflect.Int:
		Println("this type is Int")
	case reflect.Struct:
		Println("this type is Struct")
		for i := 0; i < v.NumField(); i++ {
			value := v.Field(i)
			filed_name := t.Field(i)
			Printf("value name is %s,value type is %s,value is%v\n", filed_name.Name, filed_name.Type, value.Interface())
		}
	case reflect.Float64:
		Println("this type is Float64")
	default:
		Println("I don't know what's this!")
	}
}
