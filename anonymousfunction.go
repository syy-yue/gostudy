package main
import "fmt"
func main(){
   fun1()
   fun2 := fun1
   fun2()  

   //匿名函数
   func(){
    fmt.Println("这是一个匿名函数")
   } () //匿名函数的调用，直接在这个位置加（）
   fun3 := func (){
    fmt.Println("这也是一个匿名函数")
   }
   fun3()
   fun3()
   fun3()
   //带参数的匿名函数
   func(a,b int){    //注意此处不要写成（a,int b）
     fmt.Println(a,b)
   }(1,2)
   //带返回值的匿名函数
   res1 := func(a,b int)int{
   return a+b
   }(10,20)   //如果后面没有赋值，res1的值是匿名函数的地址
   fmt.Println(res1)
}
func fun1(){
  fmt.Println("这是fun1函数")
}
