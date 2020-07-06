package main
import "fmt"

func main(){
        //实现切片的深拷贝

        //方法一
        s1 := [] int {1,2,3,4}
        s2 :=  make([] int,0)  //len:0,cap:0
        for i :=0;i<len(s1);i++{
              s2=append(s2,s1[i])
         }
        fmt.Println(s1)
        fmt.Println(s2)
    
        s1[0] = 100
        fmt.Println(s1)
        fmt.Println(s2)

        //方法二：内置函数copy
        s3 := [] int {7,8,9}
        fmt.Println(s2)
        fmt.Println(s3)

       // copy(s2,s3)      //将s3中的元素拷贝到s2中i
        copy(s3[1:],s2[2:])    //将s2中的部分拷贝到s3的部分
        fmt.Println(s2)
        fmt.Println(s3)
 }
