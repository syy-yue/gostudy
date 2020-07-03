package main
import "fmt"
func main (){
    //方法一
    s1 := Student{Person{"Alice",20},"南京工程学院"}
    fmt.Println(s1)
    //方法二
    s2 := Student{Person:Person{name:"syy",age:11},school:"南京大学"}
    //不能混搭，前面有name:后面没有school
    fmt.Println(s2)
    //方法三
    s3 := Student{}
    s3.Person.name = "宋悦悦"
    s3.Person.age = 10
    s3.school = "南京工程学院"
    fmt.Println(s3)

}
type Person struct{
    name string
    age  int
}
type Student struct{
    Person
    school string
}



