package main

import("fmt"
		"runtime"
		"time"
)
func init(){
	//获取逻辑CPU的数量
	fmt.Println("逻辑CPU的数量",runtime.NumCPU())
	
	//设置go执行的最大的cpu的数量：[1,256]
	n := runtime.GOMAXPROCS(8)
    fmt.Println(n)
}
func main(){
/*
	//获取goroot目录
	fmt.Println("GOROOT--->",runtime.GOROOT())
	//获取操作系统
 	fmt.Println("os/platform--->",runtime.GOOS)
	
    //gosched
	go func(){
		for i := 0;i < 5;i++{
			fmt.Println("goroutine...")
		}
	}()

	for i := 0; i<4;i++{
		//让出时间片，让别的goroutine执行
		runtime.Gosched()
		fmt.Println("main....")
*/
	//创建一个goroutine
	go func(){
		fmt.Println("goroutine开始....")
		//调用fun
		fun()
		fmt.Println("goroutine结束.....")
}()

	//睡一会
	time.Sleep(3*time.Second)
}

func fun(){
	defer  fmt.Println("defer....")
	//return //终止函数
    runtime.Goexit() //终止当前的goroutine  

	fmt.Println("fun函数.....")
}

