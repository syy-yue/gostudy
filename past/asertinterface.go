package main
import ("fmt"
        "math"
)
func main(){
   var t1 Triangle =Triangle{1,2,3}
   fmt.Println(t1.peri())
   fmt.Println(t1.area())
   fmt.Println(t1.a,t1.b,t1.c)

   var c1 Circle = Circle{4}
   fmt.Println(c1.peri())
   fmt.Println(c1.area())
   fmt.Println(c1.radius)

   var s1 Shape
   s1=t1
   fmt.Println(s1.peri())
   fmt.Println(s1.area())
//   fmt.Println(s1.a)   不能访问a

    var s2 Shape
    s2 = c1
    fmt.Println(s2.peri())
    fmt.Println(s2.area())

    testShape(s1)
    testShape(t1)
    testShape(c1)

    getType(t1)
    getType(c1)
    getType(s1)
//    getType(100) 
 
    var t2 * Triangle = & Triangle{3,4,20}  //注意*在定义的变量值的后面
    fmt.Println("t2:%T,%p\n,%p\n",t2,&t2,t2) 
    getType(t2)      //没加第二个else if ，打印的是布吉岛
    //加第二个else if发现t2,s，ins地址都不一样，结构体地址相同
    getType(*t2)     //打印是三角形

    getType2(t2)    
    getType2(*t2)
}
type Shape interface{
      peri() float64       //有返回值的话此处一定要注意写上
      area() float64
}
type Triangle struct{
      a,b,c float64
}
func (t Triangle) peri() float64{
      return t.a+t.b+t.c
}
func (t Triangle) area() float64{
      p:= t.peri()
      return math.Sqrt(p*(p-t.a)*(p-t.b)*(p-t.c))
}
type Circle struct{
      radius float64
}
func (c Circle) peri() float64{
      return c.radius*2*math.Pi
}
func (c Circle) area() float64{
      return  math.Pi*math.Pow(c.radius,2)
}
func  testShape(s Shape){
      fmt.Printf("周长：%.2f,面积：%.2f\n",s.peri(),s.area())  
}
func getType(s Shape){      //t1是结构体，值传递，s的值为t1的值复制了一份给他
        if ins,ok := s.(Triangle) ; ok {
           fmt.Println("是三角形,三边是",ins.a,ins.b,ins.c)
      } else if ins,ok := s.(Circle) ; ok {              //注意else要在if花括号之后，如果在下一行会报错
           fmt.Println("是圆形,半径是",ins.radius)
      }else if ins,ok := s.(*Triangle) ; ok{
           fmt.Println("ins:%T,%p\n,%p\n",ins,&ins,ins)
           fmt.Println("s:%T,%p\n",s,&s)
      } else {
           fmt.Println("布吉岛") 
     }
}
func getType2(s Shape){
     switch ins := s.(type){
     case Triangle:
           fmt.Println("是三角形",ins)

     }
}
