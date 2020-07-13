package main

import(
	"fmt"
	"reflect"
)

type Student struct{
	Name string
	Age int
	School string
}

func main(){
	/*
	struct里面的值怎么修改
	*/
	s1 := Student{"懒羊羊",5,"羊村小学"}

	//通过反射，修改对象的值，前提也是数据可以被更改
	fmt.Printf("%T\n",s1)
	p1 := &s1
	fmt.Printf("%T\n",p1)
	fmt.Println(s1.Name)
	fmt.Println((*p1).Name,p1.Name)

	//改变数值
	value := reflect.ValueOf(&s1)
	if value.Kind() == reflect.Ptr{
		newValue := value.Elem()
		fmt.Println(newValue.CanSet())

		f1 := newValue.FieldByName("Name") //也可以通过下标
		f1.SetString("宋悦悦")
		f3 := newValue.FieldByName("School")
		f3.SetString("幼儿园")
		fmt.Println(s1)
	}
}
