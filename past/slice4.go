ackage main
import "fmt"
//import "unsafe"
/*
import (
"fmt"
"unsafe"
"reflect"
)
*/
func main(){
           //a := []int{}
		   var b interface{} = []int{} //当接口类型为空的时候，可以存储任意类型的变量
		   fmt.Printf("-----%T------\n",b)
        //  fmt.Println( unsafe.Sizeof(a) )
		  //fmt.Println(a.len)
	      //hdr :=(*reflect.StringHeader)(unsafe.Pointer(&a))
		  //fmt.Println(hdr.Len)
}

