package main
import "fmt"
func main(){
    s1 := Student{"小明",20}
    fmt.Println(s1)
    //下列定义的匿名函数和匿名结构体只能使用一次
    //定义匿名函数
    func (){
    fmt.Println("匿名函数")
     }()
    //定义匿名结构体
    s2 := struct{
    name string
    age int
    }{
    name :"小红",  //注意此行和下一行都需要,
    age:20,
    }
    fmt.Println(s2)   
    
    s3 :=Worker {"小绿",20} 
    fmt.Println(s3)
}
//定义有名结构体
type Student struct{
      name  string
      age    int
}
//定义匿名字段的结构体
type Worker struct{
      string   //这种不允许再定义一个string
      int
}
