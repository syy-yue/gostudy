
package main
import "fmt"
func main(){
    var c1 Cat   //此时c1能访问A,B,的函数
    c1.test1()
    c1.test2()
    c1.test3()

    var c2 A = c1
    c2.test1() //此时只能访问test1（）
  
    var c3 C =  c1
    c3.test1()   
}
type A interface{
     test1()
}
type B interface{
     test2()
     C
}
type C interface{
     A
     B
     test3()
}
type Cat struct{
}
func (c Cat) test1(){
     fmt.Println("这是test1")
}

func (c Cat) test2(){
     fmt.Println("这是test2")
}
func (c Cat) test3(){
     fmt.Println("这是test3")
}
