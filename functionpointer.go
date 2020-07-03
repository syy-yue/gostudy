package main
import "fmt"
func main(){
  var a func()  //函数也是一钟数据类型，所以可以定义变量
  a=fun1  //将函数结构体的地址赋值给a
  a()   //a也可以调用函数
  arr2 := fun2()
  fmt.Printf("arr的类型：%T	地址:%p 数值：%d",arr2,&arr2,arr2)
  //注意此处数组的取地址也要加上& 
  arr4 := fun3()
  fmt.Printf("arr的类型：%T	地址:%p 数值：%d",arr4,&arr4,arr4)
}
func fun1(){
  fmt.Println("这是一个函数")
}

func fun2()[4]int{
//    arr1 := {1,2,3,4}  这样写是错的
      arr1 := [4] int {1,2,3,4}
    return arr1 
}
func fun3() *[4] int {
      arr3 := [4] int{1,2,3,4}
      return &arr3
}
