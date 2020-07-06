package main 
import( "fmt"
        "errors"
        "os"
        "net"
        "path/filepath"
)
func main(){
    //1  创建一个error数据
    err1 := errors.New("自己创建一个错误")
    fmt.Printf("%T\n",err1)
    fmt.Println(err1)
//    fmt.Println(*err1)  err1本身就是一个指针类型的
   //2   创建一error数据
    err2 := fmt.Errorf("错误的信息码%d",100)
    fmt.Printf("%T\n",err2)
    fmt.Println(err2)
    
    err3:= checkAge(1)
    fmt.Printf("%T\n",err3)
    fmt.Println(err3)
   // checkAge(1)   这样写是不打印东西的，因为没有返回值

   //3 查看错误的具体信息
        //1）.通过字段获取
    f, err := os.Open("/test.txt")
    if err != nil{
      fmt.Println(err)
      if ins,ok := err.(*os.PathError);ok{
         fmt.Println("1:Op:",ins.Op)
         fmt.Println("2:Path",ins.Path)
         fmt.Println("3:Err",ins.Err)
      }
      return
    }
      fmt.Println(f.Name(),"打开文件成功")
          //2）.通过函数返回
            //不知道为啥没有反映
      addr,err := net.LookupHost("today is a good day")
      fmt.Println(err)
      fmt.Println(addr)
          //3).通过模式匹配
            files,err := filepath.Glob("[")
            if err != nil && err == filepath.ErrBadPattern{
             fmt.Println(err)
             return 
      }
      fmt.Println("files:",files)

           
}
//设计一个函数
func checkAge(a int) error{
     if a<10{
     return errors.New("年龄小于0") 
     err := fmt.Errorf("年龄小于10")
     return err
     }
     fmt.Println("age is",a)
     return nil 
}
