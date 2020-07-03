package  main
import( "fmt"
        "math"
      )
func main(){
       radius :=- 3.0
       area,err := circleArea(radius)
       if err !=  nil{
       fmt.Println(err)
            if err,ok := err.(*areaError) ; ok{
              fmt.Printf("半径是：%.2f\n",err,radius)
        }
       return
       }     
       fmt.Println("圆形的面积是",area) 
       s1 := &areaError{"送月夜",radius}
       fmt.Println(s1)
}
//1.定义一个结构体，表示错误的类型
type areaError struct{
     msg string
     radius float64
}
//2.实现error接口，表示错误的类型
func (e *areaError) Error()  string{//结构体名字直接打印的话是地址
    return fmt.Sprintf("半径：%.2f,%s",(*e).radius,(*e).msg)
 }
func circleArea(radius float64) (float64,error){
     if radius < 0{
        return 0,&areaError{"半径是非法的",radius}  //结构体赋值用花括号
      }
      return math.Pi * radius *radius,nil
}
