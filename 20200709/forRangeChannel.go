package main

import("fmt"
		"time" 
)

func main(){
	/*
	通过range访问通道
	*/
	ch1 := make(chan int)
	go sendData(ch1)
	for v := range ch1{
		fmt.Println("读取数据",v)
	}
	fmt.Println("main...over...")
}

func sendData(ch1 chan int){
	for i:=0;i<10;i++{
		time.Sleep(1*time.Second)
		ch1 <- i
	}
	close(ch1) //如果没有这一句会发生死锁，因为主函数还在一直等着读取
}
