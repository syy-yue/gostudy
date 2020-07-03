package main
import "fmt"
import "strings"
func main (){
     s1 := "helloworld"
     //1.是否包含指定的内容，返回布尔值
     fmt.Println(strings.Contains(s1,"abc"))
     fmt.Println(strings.Contains(s1,"lo"))
     //2.是否包含字符串中的任意一个字符
     fmt.Println(strings.ContainsAny(s1,"abcd"))
     //3.统计指定字符或字符串出现的次数
     fmt.Println(strings.Count(s1,"llo"))
     //4.判断字符串是不是以指定字符开头或结尾
     s2 := "20200617.txt"
     if strings.HasPrefix(s2,"2020"){
     fmt.Println("2020")
     }
     if strings.HasSuffix(s2,".txt"){
     fmt.Println(".txt")
     }
     //5.查看指定字符在字符串中出出现的第一次的位置，从0开始，没有返回-1
     fmt.Println(strings.Index(s1,"h"))
     //6.查看指定字符串出现的任一指定字符第一次出现的位置
     fmt.Println(strings.IndexAny(s1,"abcd"))
     //7.查找指定字符最后一次出现的位置
     fmt.Println(strings.LastIndex(s1,"l"))
     //8.字符串的拼接,把切片的内容通过后面的指定符号拼接在一起
//     slice1 := [] string {"song","yue","yue"}
//     s2=(slice1,"*")  //执行的结果为song*yue*yue,而不是20200617.txt*song  
  //   fmt.Println(s2)
     //9.切割指定的字符串
     s4 := "13,248,3405"
     ss4 := strings.Split(s4,",")
     fmt.Println(ss4)
     for i :=0 ;i<len(ss4);i++{
     fmt.Println(ss4[i])
      }
      //10.	自己拼接自己count次
      s5 := strings.Repeat("hello",5)
      fmt.Println(s5)
       //11.替换,被操作的字符串，被替换的字符，替换字符，字符被替换几次（-1表示都被替换）
      s6 := strings.Replace(s1,"l","*",2)
      fmt.Println(s6)
      //11将字符串转为大写或者小写 
      s7  := "HelloWorld**123.."
      fmt.Println(strings.ToUpper(s7))
      fmt.Println(strings.ToLower(s7))
      //12截取子串，其他语言的有substring，go没有
      //go实现用str[start;end] 但是左闭右开
      fmt.Println(s1)
      s8:= s1[0:4]
      fmt.Println(s8)

}
