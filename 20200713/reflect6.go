package main

import(
	"fmt"
	"reflect"
	"strconv"
)

func main(){
	/*
	函数的反射:
	思路：函数也可以将看成接口变量类型
	step1 : 函数-->反射对象,value
	step2 ：kind --> func
	step3 ：Call()
	*/

	f1 := fun1
	value := reflect.ValueOf(f1)
	fmt.Printf("kind:%s,type:%s\n",value.Kind(),value.Type())

	value2 := reflect.ValueOf(fun2)
	fmt.Println("kind:%s,type:%s\n",value2.Kind(),value2.Type())

	value3 := reflect.ValueOf(fun3)
	fmt.Println("kind:%s,type:%s\n",value3.Kind(),value3.Type())

	//通过反射调用函数
	value.Call(nil)
	value2.Call([]reflect.Value{reflect.ValueOf(1000),reflect.ValueOf("syy")})
	resultValue := value3.Call([]reflect.Value{reflect.ValueOf(20),reflect.ValueOf("sss")})
	fmt.Printf("%T\n",resultValue)
	fmt.Println(len(resultValue))  //一个返回值
	fmt.Println("kind:%s,type:%s\n",resultValue[0].Kind(),resultValue[0].Type())

	s := resultValue[0].Interface().(string)   //把反射对象变成原来的类型
	fmt.Println(s)
	fmt.Printf("%T\n",s)
}

func fun1(){
	fmt.Println("fun1,have args")
}

func fun2(i int,s string){
	fmt.Println(i,s)
}

func fun3(i int,s string)(string){
	fmt.Println("fun3,have args,have return")
	return s + strconv.Itoa(i)
}
