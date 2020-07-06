package main
import "fmt"
func main(){
      s1 := Student{name:"小红"}
      fmt.Printf(s1.name)    //s1有name字段
      s1.study()     
      s1.play()
      t1 := Teacher{name:"老红"}
      fmt.Printf(t1.name)
      t1.study()

      testInterface(s1)
       
      var p1 Person
      p1 = s1
//    fmt.Printf(p1.name)   p1没有name字段，接口不能访问实现类当中的属性
      p1.study()  

      var arr [3]Person
      arr[1] = s1
      arr[2] = t1
      fmt.Println(arr)
}
type  Person interface{
     study()     
} 
type Student struct{
     name string
}
type Teacher struct{
     name string
}
//func  study(s Student){
 //    fmt.Printf("student study")
//}         注意此处和函数写法的区别
func   (s Student) study(){
     fmt.Printf("student study")
}
func  (t Teacher) study(){
     fmt.Printf("teacher study")
}
func testInterface(p Person){
     p.study()
}
func (s Student) play(){
     fmt.Printf("愉快的玩耍")
}
