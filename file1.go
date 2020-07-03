package main 

import(
		"fmt"
		"path/filepath"
        "path"
        "os"
)

func main(){
		/*
         文件操作
		*/
		//1.路径
		fileName1 := "/home/songyueyue/songyueyue.txt"
        fileName2 := "abcd.txt"

        fmt.Println(filepath.IsAbs(fileName1))  //判断是否是绝对路径，true
        fmt.Println(filepath.IsAbs(fileName2))  
        fmt.Println(filepath.Abs(fileName1))  //获取绝对路径
        fmt.Println(filepath.Abs(fileName1))  //注意filepath包不是直接import "filepath"
        
        fmt.Println("获取父目录",path.Join(fileName1,".."))  //获取父目录

		//2.创建目录
            //创建单层目录
       // err := os.Mkdir("/home/songyueyue/testDir4",os.ModePerm)  //ModePerm=777 
       // if err != nil{
	   //	fmt.Println("err:\n",err)
	   //	    return
       // }
       // fmt.Println("文件夹创建成功")   //创建成功的前提是testDir前面的目录存在
            //创建多层目录
        err1 := os.MkdirAll("/home/songyueyue/aa/bb/cc/dd/ee",os.ModePerm)  //ModePerm=777 
        if err1 != nil{
			fmt.Println("err1:\n",err1)
		    return
        }
        fmt.Println("创建多层目录成功")  //前面的aa，bb，cc可不存在

        //3.创建文件
            //创建绝对路径的文件
        file1,err2 := os.Create("/home/songyueyue/createdFile.txt")
               //如果文件存在则置为空,默认权限066,任何人可读写不可执行
        if err2 != nil{
		   fmt.Println("err2:",err2)
           return     //return不要忘写
		}
        fmt.Println(file1)
            //创建相对路径的文件
        file2,err3 := os.Create("bb.txt")  //如果文件存在则置为空
        if err3 != nil{
		   fmt.Println("err3:",err3)
           return     //return不要忘写
		}
        fmt.Println(file2)

        //4.打开文件(让当前的程序和指定文件之间建立一个链接)
        file3,err4 := os.Open("bb.txt")  //只能读文件
        if err4 != nil{
		   fmt.Println("err4:",err4)
           return     
		}
        fmt.Println(file3)
 
        file4,err5 := os.OpenFile(fileName1,os.O_RDONLY|os.O_WRONLY,os.ModePerm)
               //如果文件不存在，设置新建的文件是什么权限
        if err5 != nil{
		   fmt.Println("err5:",err5)
           return     
		}
        fmt.Println(file4)

		//5.关闭文件（程序和文件的链接断开）
        file4.Close()   
        
        //6.删除文件或者文件夹
               //删除文件非不菲空都可以，目录非空删除不了，除非用RemoveAll
        err6 :=  os.Remove("/home/songyueyue/aa/bb/cc/dd/ee")   
        if err6 !=nil {
		   fmt.Println("err6",err6)
           return
		}
        fmt.Println("删除文件夹成功")
         
}  


