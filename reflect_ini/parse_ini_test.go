package main
/*
	单元测试：
		执行go test -v：执行包内所有_test.go文件
		go test -run 函数名称
*/
import (
	"testing"
)

func TestParseIni(t *testing.T) {
	var configIni ConfigInfo
	str := ReadFile("conf.ini")
	err := ParseIni(str, &configIni)
	if err != nil {
		t.Fatal(err)
	}
}
/*
	压力测试：
		执行 go test -bench 函数名
	获取函数执行消耗时间

*/
func Benchmark_ParseIni(b *testing.B)  {
	var configIni ConfigInfo
	str := ReadFile("conf.ini")
	_ = ParseIni(str, &configIni)
}