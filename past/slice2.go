package main
import  "fmt"

func main(){
        a := [10] int{1,2,3,4,5,6,7,8,9,10}
        fmt.Println("-------1.已有数组创建切片-------")
        s1 := a[:5]  //1-4 左闭右开,从0开始数
        s2 := a[3:8]
        s3 := a[5:]
        s4 := a[:]
        fmt.Println("a",a)
        fmt.Println("s1",s1)
        fmt.Println("s2",s2)
        fmt.Println("s3",s3)
        fmt.Println("s4",s4)
        
        fmt.Printf("%p\n",&a)    //注意此处取数组的地址需要&，但是取切片地址不需要
        fmt.Printf("%p\n",s1)     //注意要打印地址的时候这里是Printf，不是Println
        fmt.Printf("%p\n",s2)
        fmt.Printf("%p\n",s3)   
        fmt.Printf("%p\n",s4)
      
        fmt.Println("------2.长度和度量-------")
        fmt.Printf("a  len:%d,cap:%d\n",len(a),cap(a))   
        fmt.Printf("s1 len:%d,cap:%d\n",len(s1),cap(s1))   //长度是实际数值的长度，容量是从开始到结束的值
        fmt.Printf("s2 len:%d,cap:%d\n",len(s2),cap(s2))
        fmt.Printf("s3 len:%d,cap:%d\n",len(s3),cap(s3))
        fmt.Printf("s4 len:%d,cap:%d\n",len(s4),cap(s4))
 
        fmt.Println("------3.更改数组的内容-----") //因为切片是引用类型，指向了底层数组
        a[4]=10
        fmt.Println(a)       
        fmt.Println(s1)
        fmt.Println(s2)
        fmt.Println(s3)
        fmt.Println(s4)

        fmt.Println("------4.更改切片的内容------")//操作切片也是操作了底层数组
        s1[2]=200
        fmt.Println(a)
        fmt.Println(s1)
        fmt.Println(s2)
        fmt.Println(s3)
        fmt.Println(s4)

        fmt.Println("------4.更改切片内容------")//不扩容的情况下其他数组和切片也改变
        s1=append(s1,1,1,1,1)
        fmt.Println(a)
        fmt.Println(s1)
        fmt.Println(s2)
        fmt.Println(s3)
        fmt.Println(s4)
        
        fmt.Println("-----5.添加元素切片扩容------")//扩容之后s1重新指向了一个新的底层数组，把原来的值复制进去，所以其他的数组和值不改变内容，并且指向的地址也不一样了
        fmt.Printf("len:%d,cap:%d",len(s1),cap(s1))
        s1=append(s1,2,2,2,2,2,2)
        fmt.Println(a)
        fmt.Println(s1)
        fmt.Println(s2)
        fmt.Println(s3)
        fmt.Println(s4)
        fmt.Printf("len：%d,cap:%d",len(s1),cap(s1))
        fmt.Printf("%p\n",&a)
        fmt.Printf("%p\n",s1)
}

        
