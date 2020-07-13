package main

import(
	"fmt"
	"reflect"
)

func main(){
	/*
	reflect对象设置实际变量的值
	*/

	var num float64 = 1.23
	fmt.Println("num的数值是：",num)

	//需要操作指针
	//通过reflect.ValueOf()获取num的value对象
	pointer := reflect.ValueOf(&num) //注意参数必须是指针才能修改其值
	newValue := pointer.Elem()

	fmt.Println("类型：",newValue.Type())
	fmt.Println("是否可以修改数据",newValue.CanSet())

	//重新赋值
	newValue.SetFloat(3.14)
	fmt.Println(num)

}
