package main
import "fmt"
func main(){
   //方法一
   b1 := Book{}
   b1.bookName = "西游记"
   b1.price =10
   
   s1 := student{}
   s1.name = "派大星"
   s1.age = 10
   s1.book = b1
   fmt.Println(s1) //输出花括号里面有花括号表示结构体嵌套
   fmt.Println(s1.book.price)  //访问嵌套结构体里面的内容
   //方法二
   s2 := student{name:"蟹老板",age:20,book:Book{bookName:"go",price:20}}
   fmt.Println(s2)
   
   s3 := student{
       name:"蟹老板",
        age:20,
        book:Book{bookName:"go",price:20},//注意这样写在末尾也需要加逗号
   }
   fmt.Println(s3)
  
   b3 := Book{bookName:"傻瓜怎么过完这一生",price:10}
   s4 := student1{name:"章鱼哥",age:20,book:&b3}
   fmt.Println(s4)
   s4.book.price=200
   fmt.Println(s4.book)
}

type Book struct{
   bookName string
   price int
}
type student struct{
   name string
   age int
   book  Book
}
type student1 struct{
    name string
    age int
    book *Book
}
