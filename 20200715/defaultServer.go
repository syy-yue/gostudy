package main

import(
	"fmt"
	"net/http"
)

//创建处理器函数
func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w,"Hello World",r.URL.Path)
}

func main(){
	/*
	调用默认的服务器
	*/

	//处理请求，映射地址
	http.HandleFunc("/",handler)	//访问根目录的时候可以调用处理器函数
									//自动转换成ServerHTTP函数，调用默认服务器
	//创建路由
	http.ListenAndServe(":8080",nil)	//监听的端口号，默认的多路复用
}
