package main
import "fmt"

func main(){
      //1.定义字符串
      s1 := "hello中国"
      //获取字符串的长度，返回的是字节个数
      fmt.Println(len(s1))
      //2.获取某个字节
      fmt.Println(s1[0])    //显示的是编码，不是字符
      a :='h'
      b := 104
      fmt.Printf("%c,%c,%c\n",s1[0],a,b)
      //3.字符串的遍历
        //通过for下标
        //通过for range
       for i :=0 ;i<len(s1)-1;i++   {
        fmt.Println("%c\t",s1[i])
          }
       for k,v := range s1{
         fmt.Println(k,v)  
        }
       
         for _,v := range s1{
          fmt.Println("%c",v)
         }

         //4.字符串是字节的集合
             //切片转成字符串
         slice1 := [] byte{65,66,67,68,69}
         s3 := string(slice1)
         fmt.Println(slice1)
         fmt.Println(s3)
             //字符串转成切片
         s4 := "abcde"
         slice2 := [] byte(s4)
         fmt.Println(slice2)  //打印的是编码值
         
         //5.字符串是不允许修改的
         fmt.Println(s4)
       //  s4[2] = 'B'  此处会报错
         fmt.Println(s4)

         
}
 
