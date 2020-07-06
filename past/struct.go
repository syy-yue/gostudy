package main
import "fmt"
func main(){
  //方法一
   var p1 Person
   fmt.Println(p1)
   p1.name =" 宋悦悦"
   p1.age = 60
   fmt.Printf("姓名：%s，年龄：%d",p1.name,p1.age)  
  //方法二
    p2 := Person {}
    p2.name = "宋小花"
    p2.age = 30
    fmt.Printf("姓名：%s，年龄：%d",p2.name,p2.age) 
  //方法三
    p3 := Person {name :"如花" ,age : 40}
    fmt.Println(p3)
  //方法四
    p4 := Person{"海绵宝宝",20}
    fmt.Println(p4)
  //除了第四种其他可以值不写全

    fmt.Printf("%T,%p",p1,&p1)
}

//1 定义结构体
type Person struct{
    name string
    age int
}
