package main
import "fmt"
func main(){
     res := fun1(12)
     fmt.Println(res) 
     fmt.Println(fun1)
}
func fun1(n int)int{
  if n==1||n==2{
   return 1 
   }
   return fun1(n-1)+fun1(n-2)
}
