package main 
import "fmt"
func main () {
 
   //1 求和 
    getSum()
    getSum(1,2,3,4)
   //2 切片
    s1 := [] int {1,2,3,4,5,6}
    getSum(s1...)    //注意此处要算切片中数值的和要...
}
 func getSum(nums ... int){
      sum := 0
      for i := 0;i < len(nums);i++{
      sum += nums[i]
  }
  fmt.Println("总和是：",sum)
}
//func func1 (s1,s2 String,num ... float)
