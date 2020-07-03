package main
import ("fmt"
        "strconv"
)
func main(){
  //1 定义一个新类型
  var m1 myint
  m1=100
  var i1 int
  i1=300
  fmt.Printf("%T,%T",m1,i1)
// m1 =i1 类型不同，不能相互赋值

  //2 定义函数类型
    res1 := fun1()
    fmt.Println(res1(10,20))

  //3 起别名
    var i3 myint2
    i3=i2
    fmt.Println(i3)
 
}
type myint int
type myfun func(int,int) string   //注意此处是func,不能是其他值
 type myint2 = int //不是重新定义数据类型

func  fun1() myfun{
    fun := func(a,b int) string{   //注意赋值符号之后为func不是fun
      s := strconv.Itoa(a)+strconv.Itoa(b)
      return s
      }
      return fun
}
