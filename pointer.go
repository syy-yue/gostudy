package main
import "fmt"

func main(){
     a := 10
     var  p1 *int
     fmt.Println(p1) //这里不能用Printf,
     p1 = &a   //让p1指向a
     fmt.Printf("%d\n",p1)  //p1的值
     fmt.Printf("%d\n",&p1)  //p1的地址
     fmt.Printf("%d\n",*p1)  //获取指针指向的变量的数值    

}
