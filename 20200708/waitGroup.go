package main

import( "fmt"
		"sync"
)

var wg sync.WaitGroup
func main(){
	/*
    waitGroup：同步等待组（比较低级,推荐通道）
		Add(),设置等待组中要执行的子goroutine的数值,增加计数器的值一定要在wait之前，减少计数器的值在哪里都可以
		Done(),给waitGroup组中的Counter数值减1,同wg.Add(-1)
		Wait(),让主goroutine处于等待
	*/
	wg.Add(2) //观察执行结果，不是一种结果  
			//	为3的时候显示死锁，因为count的值没有办法为0，main goroutine一直无法执行
              
    go fun1()
    go fun2()

	fmt.Println("main进入阻塞状态。。。等待wg中的子goroutine结束。。。")
	wg.Wait()  //表示main goroutine进入等待状态，意味这阻塞	
	fmt.Println("main...解除阻塞")
}
func fun1(){
	for i := 1;i<10;i++{
		fmt.Println("fun1()函数打印。。。A",i)
	}
	wg.Done()
}
func fun2(){
	defer wg.Done()
	for i := 1;i<10;i++{
		fmt.Println("\tfun2()函数打印。。。A",i)
	}
}
