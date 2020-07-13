package main 

import("fmt"
)

func main(){
	/*
	单向：
		chan <-	T,只支持写
		<- chan T ,只读
	定向通道：
		创建的时候往往创建双向的
		只不过在函数中限制方向，只读，只写

	*/
	ch1 :=  make(chan int) //双向，读和写
    //ch2 :=  make(chan <- int) //单向，只能读
	//ch3 :=  make(<- chan int) //单向，只能写

    //ch1 <- 100
	//data := <- ch1

	//ch2 := <- 100
	//fmt.Println(data)

	//go fun1(ch1) //ch1本身为双向通道，在函数内部只能写，在函数外部还能读
					//go fun1(ch3)在函数外只能写，不能读 
	//data := <- ch1
    
    //ch1 <- 100 //把32行注释掉，这一行保留会有死锁的错误
	go fun2(ch1)
	ch1 <- 100
	fmt.Println("main...over...")

	
}

func fun1(ch chan <- int){
    //在函数内部，对于ch1通道，只能写数据，不能读数据
	ch <- 100
	fmt.Println("fun1函数结束")
}

func fun2(ch <- chan int){
	//只读
	data := <- ch
	fmt.Println(data)
}
