package main
import "fmt"

func main(){
            //数组
            arr := [4] int {1,2,3,4}
            fmt.Println(arr)
           
             //切片
             var s1 [] int
             fmt.Println(s1)

             s2 := [] int{1,2,3,}
             fmt.Println(s2)
             fmt.Printf("%T,%T\n",arr,s2)

             s3 := make([] int ,3,8)
             fmt.Println(s3)
             fmt.Printf("容量：%d，长度：%d\n",cap(s3),len(s3))
             s3[0] = 1
             s3[1] = 2
             s3[2] = 3
             fmt.Println(s3)
//           fmt.Println(s3[3])   //通过下标操作的是切片的长度下标，不是容量下标
                   
             //append（)专门向切片的尾端添加元素 1.直接添加元素 2.添加切片   
             s4 := make([] int , 0, 5)
             fmt.Println(s4)
             s4 =append(s4,1,2)
             fmt.Println(s4)             
             s4 = append(s4,3,4,5,6,7,)  //虽然已经超过了s4的容量，但是仍然可以添加
             fmt.Println(s4)
             
             s4 = append(s4,s3...)   //添加s3的元素,后面要添加...
             fmt.Println(s4)             

             //遍历切片
             for i:=0;i<len(s4);i++{
                 fmt.Println(s4[i])
           }


             for i,v  := range s4{
             fmt.Printf("%d->%d\n",i,v)
       }
}
