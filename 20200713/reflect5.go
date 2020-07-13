package main

import(
	"fmt"
	"reflect"
)

type Person struct{
	Name string
	Age int
	Sex string
}

func (p Person) Say(msg string){
	fmt.Println("hello,",msg)
}

func (p Person) PrintInfo(){
	fmt.Println("姓名:%s,年龄：%d,性别:%s\n",p.Name,p.Age,p.Sex)
}

func (p Person) Test(i,j int,s string){
	fmt.Println(i,j,s)
}

func main(){
	/*
	通过反射来进行方法的调用
	思路：
	step1 : 接口变量-->对象反射对象：Value
	step2 : 获取对应的方法对象:MethodByName()
	step3 : 将方法对象进行调用： Call（）
	*/
	p1 := Person{"rose",19,"女"}  //结构体赋值用{}
	value := reflect.ValueOf(p1)
	fmt.Println("kind:%s,type:%s\n",value.Kind(),value.Type())

	methodValue1 := value.MethodByName("PrintInfo")
	fmt.Println("kind:%s,type:%s\n",methodValue1.Kind(),methodValue1.Type())
	
	//没有参数，进行调用
	methodValue1.Call(nil) //没有参数，直接写nil

	args1 := make([] reflect.Value,0)  //也可以传一个空的切片
	methodValue1.Call(args1)

	methodValue2 := value.MethodByName("Say")
	fmt.Println("kind:%s,type:%s\n",methodValue2.Kind(),methodValue2.Type())
	args2 := [] reflect.Value{reflect.ValueOf("反射机制")}
	methodValue2.Call(args2)

	methodValue3 := value.MethodByName("Test")
	fmt.Println("kind:%s,type:%s\n",methodValue3.Kind(),methodValue3.Type())
	args3 := [] reflect.Value{reflect.ValueOf(100),reflect.ValueOf(200),reflect.ValueOf("wonderful day")}
	methodValue3.Call(args3)

}
