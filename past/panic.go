package main

import "fmt"

//47
func main() {

    /*
        panic 和 recover

    */

    funA()
    defer myprint("defer main3......")

    funB()
    defer myprint("defer main4......")

    fmt.Println("main over...")

}

func myprint(s string) {
    fmt.Println(s)
}

func funA() {
    fmt.Println("我是第一个函数funA")
}

func funB() { //外围函数

    defer func() {
        if msg := recover();msg != nil {
            fmt.Println(msg,"程序恢复了")
        }
    }()

    fmt.Println("我是函数funB")
    defer myprint("defer funB()1.....")
    for i:=1;i<=10;i++ {
        fmt.Println("i:",i)
        //让程序中断,中断之后代码不会执行，已经执行过的defer还是会被执行的
        if i == 5 {
            panic("funB函数，恐慌")
        }
    }
    //当外围函数代码中发生了恐慌，只有其中所有的defer得函数全部执行完毕后，才会被传到外围函数的调用,
    defer myprint("defer funB()2.....")
}
