package main
import "fmt"
func main(){
   var  a interface{}
   a = 100
   fmt.Println(a)
   a = "ls"
   fmt.Println(a)
   
/*   p := phone{}    
   p = 100
   fmt.Println(p)
   p="abc"
   fmt.Println(p)
*/
    var b1 A = cat{name:"喵喵"}
    var b2 A = dog{name:"汪汪"}
    var b3 A = 100
    var b4 A = "我是女生"
 
    fmt.Println(b1)
	fmt.Println(b2)
    fmt.Println(b3)
    fmt.Println(b4)
    
    test(b1)    
    test(9099) 

    test1(b1)
    test1(999)

    //定义一个map，键为int，值为任意类型
    map1 := make(map[int]interface{})
    map1[1]=100
    map1[2]="abc"
    map1[3]=cat{name:"咪咪"}
    fmt.Println(map1[1])   
    fmt.Println(map1[2])
    fmt.Println(map1[3])
    for  i:= 1;i<4;i++{
       fmt.Println(map1[i])
  }
  
   //定义一个且切片，存储任意类型
    slice1 :=  make([] interface {},0,10)
    slice1 = append(slice1,12,3,"fg")
    fmt.Println(slice1)

    test2(slice1)
}
type A interface{
}
type  cat struct{
   name string
}
type dog struct{
   name string
}

func  test(a A){
    fmt.Println(a)
}
//匿名的空接口
func test1(a interface{}){
    fmt.Println(a)
}
func  test2(slice2 []interface{}){
     for i:=0;i<len(slice2);i++{
 //  fmt.Println(slice2[i]) 
     fmt.Printf("第%d个数据：%v\n",i+1,slice2[i]) 
 }
}
