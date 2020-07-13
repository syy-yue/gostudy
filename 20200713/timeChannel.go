package main

import("fmt"
		"time"
)

func main(){
	/*
	1.func NewTimer(d Duration) *Timer
			创建一个计时器，d时间以后触发
	*/

	timer := time.NewTimer(3*time.Second)
	fmt.Printf("%T\n",timer)
	fmt.Println(time.Now()) //此处注意是time不是timer

	//此处等待channel中的数据，会阻塞3秒
	ch2 := timer.C
	fmt.Println(<- ch2)

	//新建一个计时器
	timer2 := time.NewTimer(5*time.Second)
	//开始goroutine,来处理出发后的事件
    go func(){
		<- timer2.C
		fmt.Println("Timer 2 结束咯。。。开始。。。")
	}()
	
	time.Sleep(3*time.Second) //在定时时间内没，没执行完的时候可以取消定时
	flag := timer2.Stop()
    if flag{
		fmt.Println("Timer 2停止了。。。")
	}
}
