package main

import("fmt"
)

func main(){
	/*
	双向：
		chan T
			chan <- data, 发送数据，写出
			data <- chan, 获取数据, 读取
	单向:
			chan <- T, 只支持写
			<- chan T, 只读	
	*/
	//双向通道例子
	ch1 := make(chan string)
	done := make(chan bool)
    
	go sendData(ch1,done)	
	
	data := <- ch1
	fmt.Println("子goroutine传来数据\n",data)

	ch1 <- "主 goroutine"
	
	<- done
	fmt.Println("main...over.,.")	
	
}

func sendData(ch1 chan string,done chan bool){
	ch1 <- "子goroutine"

	data := <- ch1
	fmt.Println("main gotoutine传来数据\n",data)
	
    done  <- true
    fmt.Println(done)
}

