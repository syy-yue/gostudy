package main

import ("fmt"
        "os")

func main(){
  /*
  FileInfo : 文件信息
    type FileInfo interface {
      Name() string        //文件名.扩展名
      Size() int64         // 文件大小，字节数 12340
      Mode() FileMode      //文件权限 -rw-rw-rw-
      ModTime() time.Time  //修改时间 2018-04-13 16:30:53 +0000 CST
      IsDir()  bool        //是否文件夹
      Sys() interface{}    //基础数据源接口(can return nil)

    }
   */
    fileInfo,err := os.Stat("/home/songyueyue/songyueyue.txt")
    if err!=nil{
      fmt.Printf("err:",err) 
      return
    }
    fmt.Printf("%T\n",fileInfo)
    fmt.Println(fileInfo.Name())
    fmt.Println(fileInfo.Size())
    fmt.Println(fileInfo.Mode()) 
    fmt.Println(fileInfo.ModTime())
    fmt.Println(fileInfo.IsDir())
}
