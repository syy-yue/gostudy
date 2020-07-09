package main

import("fmt"
		"time")

func main(){
	ch1 := make(chan int)
	go func() {
		fmt.Println("子goroutine开始执行...")
//		time.Sleep(3*time.Second)
		data := <- ch1
		fmt.Println("data",data)
    }()

	time.Sleep(5*time.Second)
	ch1 <- 10   //当子goroutine去读之后才能解除阻塞
                //针对没有缓存的通道，发送和接收默认都是阻塞的
	fmt.Println("main...over...")

	ch := make(chan int)
    ch <- 10
}

