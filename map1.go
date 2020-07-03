package main
import "fmt"

func main(){
     //1.创建map
      var map1 map [int] string  //没有初始化，nil，无法直接使用
      var map2 = make(map[int]string)    
      var map3 = map [string] int {"Song":99,"yue":88}
      fmt.Println(map1)
      fmt.Println(map2)
      fmt.Println(map3)
   
      fmt.Println(map1==nil)
      fmt.Println(map2==nil)
      fmt.Println(map3==nil)
     
     //2.map nil 
     // map1[1]="hello"
       if map1 == nil {
       map1 = make (map [int] string)
       fmt.Println(map1 == nil)

      //3.赋值
       map1[0] = "hello"
       
      //4.根据key获取对应的value值，如果value不存在，则得到默认值
      //判断value的值是不是存在（防止默认值与赋值的分不清）
       fmt.Println(map1)
       fmt.Println(map1[0])
       
       v1,ok := map1[0]
       if ok{
            fmt.Println("对应的数值为：",v1)
        }else{
            fmt.Println("操作的值不存在，获取的，默认值",v1)
        }
       
      //5.修改数据
       fmt.Println(map1)
       map1[0]="world"
       fmt.Println(map1[0])
      
      //6. 删除数据，删除key就好了    
       delete(map1,0)
       fmt.Println(map1)

      // 7.长度
       fmt.Println(len(map1))

}

}

