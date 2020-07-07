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
	//bufio下的读操作
     //创建Reader对象
	b1 := bufio.NewReader(file)
	/*
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
    */
	//3.ReadString()
	s1,err1 := b1.ReadString('8')  //6,\n这些字符都可当作分隔符
	fmt.Println(err1)
	fmt.Println(s1)

	for{
		s1,err := b1.ReadString('\n')
		if err == io.EOF{
			fmt.Println("读取完毕")
			break
		}
		fmt.Println(s1)
	}

	//4. ReadBytes()
	data,err := b1.ReadBytes('\n')
	fmt.Println(err)
	fmt.Println(string(data))

	//Sacnner等待键盘输入然后回车，如果输入中间有空格的话，空格之后的不输出
	s2 := ""
	fmt.Scanln(&s2)
	fmt.Println(s2)
	
	//输入中间有空格也可以打印
	b2 := bufio.NewReader(os.Stdin)
	s2,_ := b2.ReadString('\n')
	fmt.Println(s2)

		
}
