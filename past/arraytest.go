package main
import "fmt"
func main (){
 arr := [2]int {1,2}
 arr1 := [2]*int {&1,&2}
 fmt.Println(arr1)
}
