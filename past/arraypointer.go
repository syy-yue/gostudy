package main
import "fmt"
func main(){
   //创建一个普通的数组
   arr1 := [4] int{1,2,3,4}
   fmt.Println(arr1)
   
    //创建一个指针，存储该数组的地址->数组指针
    var  p1 *[4]int
    p1 =& arr1
    fmt.Println(p1)
    fmt.Printf("%p\n",p1)    //数组p1的地址
    fmt.Printf("%p\n",&p1)   //p1的地址
    
    //通过数组指针，操作数组
   (*p1)[1]=10   //注意此处要用括号括起来
   fmt.Println(arr1)

   p1[0]=5      //简化写法（数组就是指针？）
   fmt.Println(arr1)
   
   //指针数组
    a := 1
    b := 2
    c := 3
    d := 4
    arr2 := [4] int{a,b,c,d}  //数组中存储的是a,b,c,d的数值，值传递
    arr3 := [4] * int{&a,&b,&c,&d}
    fmt.Println(arr2)
    fmt.Println(arr3)
    arr2[0]=100
    fmt.Println(arr2)
    fmt.Println(a)     //注意此时的a的值没有变化，看截图1
    //arr3=200   错误，不能这样写，arr3存的是地址
    *arr3[0] =200
    fmt.Println(arr3)
    fmt.Println(a)  //此时a的值已经改变了
    b=500
    fmt.Println(arr2)
    fmt.Println(arr3)
    for i:=0;i<len(arr3);i++{
    fmt.Println(*arr3[i])
    }    

   
  
}
