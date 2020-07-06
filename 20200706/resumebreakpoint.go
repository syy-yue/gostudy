package main

import ("fmt"
		"os"
		"io"
		"log"
	)

func main() {
	/*
	Seek(offset int64,whence int) (int64, error),设置指针光标的位置
		第一个参数：偏移量
		第二个参数：如何设置
	*/

	fileName := "/home/songyueyue/aa.txt"
    file,err := os.OpenFile(fileName,os.O_RDWR,os.ModePerm)
    if err != nil{
		log.Fatal(err)
	}
	defer file.Close()

	//读写
	bs := []byte{0}
	file.Read(bs)
	fmt.Println(string(bs))

	file.Seek(4,io.SeekStart)//SeekStart=0 SeekCurrent=1 SeekEnd=2
	file.Read(bs)
	fmt.Println(string(bs))

	file.Seek(2,0)
	file.Read(bs)
	fmt.Println(string(bs))

	file.Seek(0,io.SeekEnd)
	file.WriteString("ABC")

}
