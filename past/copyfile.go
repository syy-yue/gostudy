package main

import ("fmt"
		"os"		
		"io"
		"io/ioutil"
	)

func main () {
	/*
	拷贝文件(文字，图片)
    三种方法：
		方法一：io包下的Read()和Write()方法实现
        方法二：io包下的Copy()方法实现
        方法三：io包里的ioutil包里的ReadFile()和WriteFile()方法
	*/
	srcFile := "/home/songyueyue/aa.txt"
	destFile := "/home/songyueyue/ab1.txt"
	//total,err := CopyFile1(srcFile,destFile)
	//total,err := CopyFile2(srcFile,destFile)
	total,err := CopyFile3(srcFile,destFile)
	fmt.Println(total,err)
} 

//该函数：用于通过io操作实现文件的拷贝，返回值是拷贝的总数量（字节）,错误
func CopyFile1(srcFile,destFile string)(int,error){
	//打开两个文件
	file1,err := os.Open(srcFile)
	if err != nil{
		return 0,err	
	}
	file2,err := os.OpenFile(destFile,os.O_WRONLY|os.O_CREATE,os.ModePerm)
	if err != nil{
		return 0,err
	}

	defer file1.Close()
	defer file2.Close()

    //读写
	bs :=  make([]byte,1024,1024)
	n := -1  //读取的数据量  为什么在n赋值为1之后运行total还是12
    total := 0
	for{
		n,err = file1.Read(bs)
		if err == io.EOF || n == 0{
			fmt.Println("拷贝完毕")
			break
		}else if err != nil{
			fmt.Println("拷贝的时候出错了")
			return total,err
		}
			total += n
			file2.Write(bs[:n])
		}
	return total,nil
}

//此函数error返回nil表示成功
func CopyFile2(srcFile, destFile string)(int64,error){	
	file1,err := os.Open(srcFile)
	if err != nil{
		return 0,err	
	}
	file2,err := os.OpenFile(destFile,os.O_WRONLY|os.O_CREATE,os.ModePerm)
	if err != nil{
		return 0,err
	}
	
	defer file1.Close()
    defer file2.Close()

	return io.Copy(file2,file1)
}

//此函数返回int64会报错
func CopyFile3(srcFile,destFile string)(int,error){  
	bs,err := ioutil.ReadFile(srcFile)
	if err != nil{
	return 0,err
	}
	err = ioutil.WriteFile(destFile,bs,0777)
    if err != nil{
		return 0,err
	}
	return len(bs),nil
}
