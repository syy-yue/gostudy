package main
import "fmt"

func main(){
        arr := [5] int{66,10,15,3,7} 
//           sum :=0
//           for _, v:=range arr1{
//           sum += v
//        fmt.Println(sum)}
         for i:=0;i<4;i++{
          for j:=0;j<4-i;j++{
                if arr[j]>arr[j+1]{
                arr[j],arr[j+1]=arr[j+1],arr[j]

               // arr[j]=temp
               // temp=arr[]j
               // arr[j+1]=arr[j]

}
                          
}               
}
            fmt.Println(arr)
}
