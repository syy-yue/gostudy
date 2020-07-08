package main

import("fmt"
		"time"
		"math/rand"
		"sync"
)

var ticket int = 10 //10张票

var mutex sync.Mutex //创建锁头

var wg sync.WaitGroup  //设置同步等待组

func main() {
	/*
	4个goroutine,模拟4个售票口
    
	注意：在使用互斥锁的时候，对资源操作完，一定要解锁，否则会出现程序异常
          可以使用defer语句一起使用
	*/
    wg.Add(4)


	go 	saleTickets("售票口1")
	go 	saleTickets("售票口2")
	go 	saleTickets("售票口3")
	go 	saleTickets("售票口4")
    
    wg.Wait()
	fmt.Println("程序结束咯")
	//time.Sleep(100*time.Second) //设置睡眠时间，否则可能只执行main goroutine,但是有同步等待组的时候就不需要睡眠了
}
func saleTickets(name string){
	rand.Seed(time.Now().UnixNano())
	for{
		//上锁,在上锁开锁之间的资源每次只能被一个协程访问
		mutex.Lock()
		defer wg.Done()
		if ticket > 0{
			time.Sleep(time.Duration(rand.Intn(3000))*time.Millisecond)
			fmt.Println(name,"售出",ticket)
			ticket--
		//	wg.Done()
		}else{


0000
			mutex.Unlock()  //条件不满足，也要解锁，因为break之后下面的解锁执行不到
			fmt.Println(name,"售完，没有票了")   
			break
		}
		mutex.Unlock() //解锁
	}
}
