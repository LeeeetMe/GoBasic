package main

import (
	"IrisProduct/common"
	"IrisProduct/repository"
	"database/sql"
	"fmt"
	"reflect"
)

//获取返回值， 获取一条
func GetResultRow(rows *sql.Rows) map[string]string {
	column, err := rows.Columns()
	if err != nil {
		return nil
	}
	size := len(column)
	argsArray := make([]interface{}, size)
	tmpArray := make([][]byte, size)
	for index, _ := range argsArray {
		argsArray[index] = &tmpArray[index]
	}
	record := make(map[string]string)
	rows.Next()
	{
		err := rows.Scan(argsArray...)
		if err != nil {
			panic(err.Error())
		}
	}
	for index, value := range column {
		record[value] = string(tmpArray[index])
	}
	return record
}

func main() {
	db, err := common.NewMysqlConn()
	if err != nil {
		panic(err.Error())
	}
	rep := repository.NewProductManager("table", db)
	pro,err := rep.SearchProduct(1)
	if err != nil {
		panic(err.Error())
	}
	e := reflect.ValueOf(pro).Elem()
	fmt.Println(e)
}
