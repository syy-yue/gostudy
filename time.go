package main
import ("fmt"
        "time"
       )
func main(){
   //1。 获取当前的时间
   t1 := time.Now()
   fmt.Printf("%T\n",t1)
   fmt.Println(t1)

   //2.  获取指定的时间
   t2 := time.Date(1996,9,20,8,40,0,0,time.Local)
   fmt.Println(t2)

   //3.  time<—>string之间的互相转换
   s1 := t1.Format("2020年7月2日99")  //注意此处为t1,t2错误
   fmt.Println(s1)

   s2 := t1.Format("2002/01/02")
   fmt.Println(s2)
 
   s3 := "1988年6月6日"
   t3,err := time.Parse("2006年3月2日",s3)  //注意此处的格式不能错
   if err != nil {
        fmt.Println("err",err) 
  }
  fmt.Printf("%T\n",t3)
  fmt.Println(t3)
 
  fmt.Println(t1.String())
  
  //4   根据当前时间，获取指定的值 
  year,month,day := t1.Date()
  fmt.Println(year,month,day)
  hour,min,sec := t1.Clock()
  fmt.Println(hour,min,sec)

  year2 := t1.Year()
  fmt.Println("年",year2)
  fmt.Println(t1.YearDay())  //年过了多少天

  month2 := t1.Month()
  fmt.Println("月",month2)
  fmt.Println("日",t1.Day())
  fmt.Println("时",t1.Hour())
  fmt.Println("分",t1.Minute())
  fmt.Println("秒",t1.Second())
  fmt.Println("纳秒",t1.Nanosecond())

}
