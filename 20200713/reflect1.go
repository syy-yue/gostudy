package main

import(
	"fmt"
	"reflect"
)

func main(){
	/*
	知道原有类型情况下的转换
	*/
	var num float64 = 1.23
    //接口类型变量-->反射类型变量
	value := reflect.ValueOf(num)

    //反射类型变量-->接口类型变量
	convertValue := value.Interface().(float64)
	fmt.Println(convertValue)

	/*
	反射类型变量-->接口类型变量，理解为强制转换
	Golang对类型要求非常严格，类型一定要完全符合
    一个为*float64,一个为float64，如果弄混,会panic
	*/
	pointer := reflect.ValueOf(&num)
	convertPointer := pointer.Interface().(*float64)
	fmt.Println(convertPointer)	
}
