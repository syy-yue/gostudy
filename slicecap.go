package main
import "fmt"

func main (){
     var s1 = [] int {1,2,3}
     s1 = append(s1,4,5,6,7,8)
     fmt.Printf("%d,%d,%d\n",len(s1),cap(s1),s1)     
}

