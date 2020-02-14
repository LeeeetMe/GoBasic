package base

import (
	"fmt"
	_ "fmt"
	"reflect"
)

func Reflect_Test() {
	/*
		reflect.TypeOf：
			直接给到了我们想要的type类型，如float64、int、各种pointer、struct 等等真实的类型
		reflect.ValueOf：
			直接给到了我们想要的具体的值，如1.2345这个具体数值，或者类似&{1 "Allen.Wu" 25} 这样的结构体struct的值
	*/
	var v int = 10
	var i int = v
	ref_value := reflect.ValueOf(i)
	fmt.Printf("i value Type is %T value is %v\n", ref_value, ref_value)
	// 这里注意是传入的值为指针类型&i，不是直接为&ref_value
	ref_pointer := reflect.ValueOf(&i)
	fmt.Printf("i pointer Type is %T, point value is %v\n", ref_pointer, ref_pointer)
	ref_type := reflect.TypeOf(i)
	fmt.Printf("i Type is %T, Type value is %v\n", ref_type, ref_type)

	/*
		将通过reflect.ValueOf获取到的值转换为可以使用的值
		value.Interface().(类型)
	*/

	converPointer := ref_pointer.Interface().(*int)
	fmt.Printf("value is %d type is %T\n", converPointer, converPointer)
	converValue := ref_value.Interface().(int)
	fmt.Printf("value is %d type is %T\n", converValue, converValue)
	// 传入参数后获取类型
	// p := user{"alex", 22, 100.0}
	// getAnything(p)
	// setValues()
	//reflectCallFunc()
}

type user struct {
	name  string
	age   int
	score float64
}

func (u user) GetName() {
	fmt.Println("My Name is ", u.name)
}
func (u *user) SetCurName(name string) {
	u.name = name
	fmt.Println("My Name is ", u.name)
}

func getAnything(input interface{}) {
	/* 获取type、value、method
	- Name() 返回类型的名称

	结构体：都是通过struct.TypeOf的返回值获取
		字段个数：NumField()
		函数个数：NumMethod()
	*/
	println()
	// 获取值
	getValue := reflect.ValueOf(input)
	fmt.Printf("that value is %#v\n", getValue)
	// 获取Type
	getType := reflect.TypeOf(input) //等同↓↓↓
	// getType := getValue.Type()
	fmt.Printf("that type is %s\n", getType.Name())

	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i)
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}
	// 函数名称首字母必须是大写，才可以计入该结构体的函数
	for j := 0; j < getType.NumMethod(); j++ {
		m := getType.Method(j)
		fmt.Printf("func name is %s: %v\n", m.Name, m.Type)
	}
}

func setValues() {

	/*
		常量赋值：
			- 需要传入的参数是* float64这个指针，然后可以通过pointer.Elem()去获取所指向的Value，注意一定要是指针。
				- 如果传入的参数不是指针，而是变量，那么
					通过Elem获取原始值对应的对象则直接panic
			- CantSet()表示是否可以重新设置其值，如果输出的是true则可修改，否则不能修改。
			- reflect.Value.Elem() 返回的是【原始值对应的反射对象】，只有原始对象才能修改，当前反射对象是不能修改的
			  也就是说如果要修改反射类型对象，其值必须是“addressable”【对应的要传入的是指针，同时要通过Elem方法获取原始值对应的反射对象】
	*/
	var f float64 = 0.21
	// 传入的值必须为指针类型，否则Elem会报错
	pointer := reflect.ValueOf(&f)
	newValue := pointer.Elem()
	fmt.Println("pointer is", pointer.Type())
	fmt.Println("newValue 是否可以赋值：", newValue.CanSet())
	newValue.SetFloat(100)
	fmt.Println("newValue is", newValue)
}

func reflectCallFunc() {
	/*
		调用函数，并传参
	*/
	p := user{
		name:  "alice",
		age:   18,
		score: 100,
	}
	p.SetCurName("alex")
	fmt.Println(p)

	p_value := reflect.ValueOf(&p)
	p_method := p_value.Elem().MethodByName("SetCurName")
	m_arguments := []reflect.Value{reflect.ValueOf("大蠢猪")}
	p_method.Call(m_arguments)
	fmt.Print(p)
}
