package  main
import "fmt"

func main() {

/*go语言支持函数式编程：
      支持将一个函数作为另一个函数的参数
      也支持将一个函数作为另一个函数的返回值
*/
        r1 := increment()
        fmt.Printf("%T\n",r1)
        fmt.Println(r1)
        
        v1 := r1()
        fmt.Println(v1)
        v2 := r1()
        fmt.Println(v2)   
        fmt.Println(r1())
        fmt.Println(r1())
 
        r2 := increment()
        fmt.Printf("%T\n",r2)
        fmt.Println(r2)

                  
                  



}

    func  increment()func() int{ //外层函数
            //定义了一个局部变量
            i :=0 ;
            //内层函数
            fun := func ()int{
              i++
              return i
           }
          return fun
}
