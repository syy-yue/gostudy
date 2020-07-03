package main
import "fmt"
import "sort"
func main(){
     map1 := make(map [int] string)
     map1[0] = "宋"
     map1[1] = "悦"
     map1[2] = "悦"
     map1[3] = "宋"
     map1[4] = "悦"
     map1[5] = "悦"
     
     //1.遍历map
     for k,v := range map1 {  //注意此处是:=而不是=
     fmt.Println(k,v)    //多运行几次发现运行结果顺序不一样
    }
      
     //2.顺序打印map的值
      keys := make([] int,0,len(map1))
      fmt.Println(keys)
      
      for k,_:= range map1{
      keys = append(keys,k)
   }     
      fmt.Println(keys)
   
      //冒泡排序，默认升序
      sort.Ints(keys) //字符串排序默认使用sort.Strings(_)


      for _,key :=  range keys{
      fmt.Println(key,map1[key])
     }

} 
