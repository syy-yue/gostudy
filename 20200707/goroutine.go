package main

import("fmt"
	   "time"
)

func main() {
	/*
	一个goroutine打印数字，一个goroutine打印字母，观察运行结果
	*/
   
    //1.先创建并启动子goroutine,执行printNum() 
    go printNum()

	//2.main中打印字母
	for i:=1;i<=100;i++{
		fmt.Printf("\t主goroutine中打印字母: A %d\n",i)

	time.Sleep(1*time.Second)  //除通道外，睡眠也能让子程序执行完再结束
	fmt.Println("main......over......")
}
	
}
func printNum(){
	for i:=1;i<=100;i++{
		fmt.Printf("子goroutine中打印数字:%d\n",i)
}
}
