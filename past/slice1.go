package main
import "fmt"

func main(){
        s1 := [] int {1,2,3}   //创建切片的时候先创建了一个底层数组，底层数组的地址赋值给s1，s1是引用类型的数据
        fmt.Println(s1)
        fmt.Printf("len:%d,cap:%d\n",len(s1),cap(s1))
        fmt.Printf("%p\n",s1)  //s1为切片的地址，不需要加&
        fmt.Printf("%p\n",&s1)  //打印值为s1本身的地址

        s2 := [] int{1,2,3}
        fmt.Println(s2)
        fmt.Printf("%d,%d\n",len(s2),cap(s2))
        fmt.Printf("%p\n",s2)

        s2 = append(s2,4,5)
        fmt.Println(s2)
        fmt.Printf("%d,%d\n",len(s2),cap(s2))
        fmt.Printf("%p\n",s2)
        
}

