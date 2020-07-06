package main

import( "os"
		"fmt"
		"bufio"
)

func main(){
	/*
	bufio:高效读和写
    	buffer缓存
		io ： inut/output
	
	将io包下的Reader，Write对象进行包装，将缓存的包装，提高读写的效率
  
		ReadBytes()
		ReadString()
		ReadLine()
	*/

		fileName := "/home/songyueyue/aa.txt"
		file,err := os.Open(fileName)
		if err != nil{
			fmt.Println(err)
			return
		}
		defer file.Close()

        //创建Reader对象
		b1 := bufio.NewReader(file)
		//1.Read(),高效读取
        p := make([]byte,10000)
		n1,err := b1.Read(p)
		fmt.Println(n1)
		fmt.Println(string(p[:n1]))

		//2.ReadLine()
		data,flag,err := b1.ReadLine()
		fmt.Println(flag)
		fmt.Println(err)
		fmt.Println(data)
		fmt.Println(string(data))
}
