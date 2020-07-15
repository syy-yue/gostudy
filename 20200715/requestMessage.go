package main

import(
	"fmt"
	"net/http"
)

//创建处理器函数
func handler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w,"测试http协议")
}

func main(){
	//调用处理器处理请求
	http.HandleFunc("/http",handler)	//调用HandleFunc，直接调用处理器函数
										//调用Handle，需要用结构体实现serveHTTP方法
	//路由
	http.ListenAndServe(":8080",nil)
}
