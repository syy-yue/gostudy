package main

import("fmt"
		"strconv"
)

func main(){
	/*
		非缓冲通道： make(chan T)
					一次发送数据，一次接收数据，都是阻塞
		缓冲通道： make(chan T,capacity)
                    发送:缓冲区的数据满了，才会阻塞
                    接收:缓冲区的数据空了, 才会阻塞
	*/
  	ch1 := make(chan int)
    fmt.Println(len(ch1),cap(ch1))
	
    	
	ch2 := make (chan int,5)
    fmt.Println(len(ch2),cap(ch2))
    ch2 <- 100
    ch2 <- 200
    ch2 <- 300
    ch2 <- 400
    ch2 <- 500
	fmt.Println(len(ch2),cap(ch2))
    //ch2 <- 600 超出容量，发生死锁

	ch3 := make(chan string,5)
	go sendData(ch3)
    for{
		v,ok := <-ch3
		if !ok{
			fmt.Println("读完了",ok)
            break
}
        fmt.Println("\t读取的数据是：",v)
	}
    fmt.Println("main...over...")

}
	
func sendData(ch chan string) {
	for i:=0;i<10;i++{
		ch <- "数据" + strconv.Itoa(i)
		fmt.Printf("子goroutine中写出第%d个数据\n",i)
	}
    close(ch)   //不要忘记关闭通道
}
