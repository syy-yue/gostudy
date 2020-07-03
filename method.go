package main
import "fmt"
func main(){
    s := Student{"小明",20}
    s.study()   //接收者类型的值来调用
     
    s2 :=& Student{"小红",30}
    fmt.Printf("%T\n",s2) //s2为指针类型
    s2.rest()
    s.rest() //此时下面方法取的是s的地址
    s2.study()
   
    c1 := Cat{"yellow",1}
    c1.rest()
}
type  Student struct{
  name string 
  age int
}
type Cat struct{
  color string
  age int
}
func (s Student)  study (){
  fmt.Println(s.name,"在学习")
}
func (s *Student) rest(){
  fmt.Println((*s).name,"在休息") 
    //注意此处要是按照指针访问的话一定要加括号
  fmt.Println(s.name,"is resting")//结构体的访问可以这样简写
}
func (s *Cat)   rest(){ //方法名可以同名，但是接收者需要不同
  fmt.Println("cat is resting")
}
