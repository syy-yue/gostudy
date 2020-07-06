ackage  main
import "bufio"
import  "fmt"
import "os"
func main(){
     //    var  x int
     //    var  y float64
     //   fmt.Println("请输入一个整数，一个浮点型")    //此处输入数字空格数字就好了
     //   fmt.Scanln(&x,&y)     //读取键盘的输入，通过操作地址，赋值给x和y
     //    fmt.Printf("a的数值：%d ,b的数值：%f\n",x,y)         
           
     //     fmt.Scanf("%d,%f",&x,&y )         //此处输入数字,数字就好了
     //    fmt.Printf( "x:%d,y:%f\n",x,y )

           fmt.Println("请输入一个字符串：")
           reader := bufio.NewReader(os.Stdin)
           s1,_ := reader.ReadString('\n')
           fmt.Println("读到的数据",s1)


}
