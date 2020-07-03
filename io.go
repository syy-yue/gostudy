package main

import (
    "fmt"
//    "io"
    "os"
)

//5
func main() {
    /*
    读取数据：
        Reader接口：
            Read(p [[byte)(n int,error)


    读取本地aa.txt文件中的数据、

    */
    //1。 打开文件
    fileName := "/home/songyueyue/aa.txt"
    file,err := os.Open(fileName)
    if err != nil {
        fmt.Println("err:",err)
        return
    }

    //3。 关闭文件
    defer file.Close()

    //2。 读取数据 abcdefghij
    bs := make([]byte,4,4)


    //第一次读取
    n,err := file.Read(bs)
    fmt.Println(err) //<nil>
    fmt.Println(n) //4
    fmt.Println(bs) //[97 98 99 100]  因为a,b,c,d是按byte读取的
    fmt.Println(string(bs)) //abcd

    //第二次读取
    n,err = file.Read(bs)
    fmt.Println(err) //<nil>
    fmt.Println(n) //4
    fmt.Println(bs) //[101 102 103 104]
    fmt.Println(string(bs)) //efgh

    //第三次读取
    n,err = file.Read(bs)
    fmt.Println(err) //<nil>
    fmt.Println(n) //2     这边不是2而是4？？？
    fmt.Println(bs) //[105 106 103 104]    输出的值是[105 106 10 10]
    fmt.Println(string(bs)) //ijgh
/*
    //第四次读取
    n,err = file.Read(bs)
    fmt.Println(err) //EOF
    fmt.Println(n) //0
    fmt.Println(bs) //[105 106 103 104]
    fmt.Println(string(bs)) //ijgh
*/

/*
   abcd
   efgh
   ij
   读取到文件末尾，结束读取操作
*/
 /*  n := -1
    for {
        n,err = file.Read(bs)
        if n==0 || err == io.EOF{
            fmt.Println("读取到文件末尾，结束读取操作")
            break
        }
        fmt.Println(string(bs[:n]))
    }

*/
}
