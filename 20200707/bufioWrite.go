package main

import ("fmt"
		"os"
		"bufio"
)

func main(){
	/*
	将io包下的Reader，Write对象进行包装，带缓存的包装，提高读写的效率
    
		func (b *Writer) Write(p [] byte) (nn int,err error)
		func (b *Writer) WriteByte(c byte) error
        func (b *Writer) WriteRune(r rune) (size int, err error)
		func (b *writer) WriteString(s String) (int,error)
	*/

	fileName := "/home/songyueyue/aa.txt"
	file , err := os.OpenFile(fileName,os.O_CREATE|os.O_WRONLY,os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}	
	defer file.Close()

	
	//情况一：写入的比缓冲区小
	w1 := bufio.NewWriter(file)
	//n,err := w1.WriteString("helloworld")
	//fmt.Println(err)
	//fmt.Println(n)
	//w1.Flush() //刷新缓冲区
	

	//情况二： 写入的比缓冲区大
	for i  := 1;i<=1000;i++{
		w1.WriteString(fmt.Sprintf("%d:hello",i))
	}
	w1.Flush()
}
