package main 
import(
         "fmt"
         "os"
 //        "log"
       )
func main(){
    f,err := os.Open("/home/songyueyue/a.txt")
    fmt.Printf("%T,%p",f,f)
    //   *f.Open("/a.txt")  不能这样使用
    if err != nil{
       fmt.Println(err)
       return    //它和第二句是第一种打印err的方法
      // log.Fatal(err) 第二种打印错误的方法
     }
    //根据f进行文件读或写
   // fmt.Println(f.Name(),"opened successfully")
      fmt.Println(f.Name(),"打开文件成功")
}
