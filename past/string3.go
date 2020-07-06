package main
import "fmt"
import "strconv"
func main(){
   //fmt.Println("aa"+100)
   //1 string到bool     strconv.ParseBool()
   s1 := "true"
   b1,err := strconv.ParseBool(s1)
   if err != nil {
      fmt.Println(err)
      return
   }   
   fmt.Printf("%T,%t\n",b1,b1)
   // 2 booi到string    strconv.Format()
   ss1 := strconv.FormatBool(b1)
   fmt.Printf("%T,%s",ss1,ss1)
  
   // 3 int  string  互转 strconv.Atoi() strconv.Itoa()
    

   


   
   
}

