package main
import "fmt"

func main(){
  fmt.Printf("%T\n",fun1)
  fmt.Println(fun1)
  var c func(int,int)  //注意此处不是fun2这种函数名
  fmt.Println(c)
  c = fun1   
 //将fun1的值（函数体的地址）赋值给c;加括号表示调用，将返回值赋值给c
  fmt.Println(c)   //显示结果与第二行相同

  fun1(10,20)
  c(10,20)   //也能向c里面传参数;但是带括号的不能这样用
}
func fun1(a,b int){
  fmt.Printf("a:%d,b:%d",a,b)
}
