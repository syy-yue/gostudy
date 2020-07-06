package main

import ("fmt"
        "os"
        "log"
)

func main(){
	/*
    写出数据(三步)
	step1 : 打开文件
	step2 : 写出数据
	step3 : 关闭文件
     */
    
	fileName := "/home/songyueyue/ab.txt"
	//打开未创建的文件
	file,err := os.OpenFile(fileName,os.O_CREATE|os.O_WRONLY,os.ModePerm) 
	//打开已经创建的文件
	// file,err := os.Open(fileName,os.O_WRONLY,os.ModePerm) 
    //在文件后面添加内容
    // file,err := os.OpenFile(fileName,os.O_CREATE|os.O_WRONLY|os.O_APPEND,os.ModePerm) 
    if err != nil{
			fmt.Println(err)
            return
        }
        
    defer file.Close()   //Close的c要大写
    
	//写出数据
    bs := []byte{65,66,67,68,69,70}  //A,B,C,D,E,F  
    n,err := file.Write(bs[:6])
    fmt.Println(n)
    HandleErr(err)

	//写出字符串(在文件末尾追加的)
    n,err = file.WriteString("Hello World")
    fmt.Println(n)
    HandleErr(err)
    file.WriteString("\n")

    n,err = file.Write([] byte("today"))
    fmt.Println(n)
    HandleErr(err)
      
}
	func HandleErr(err error){
    	if err != nil {
			log.Fatal(err)
	}
}
