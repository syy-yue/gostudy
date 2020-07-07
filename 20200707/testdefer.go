package main

import "fmt"

func main(){
	a := test1()
	fmt.Println(a)
	b := test2()
	fmt.Println(b)
}

func test1() (k string){
	k = "hhhh"
    defer func(){
		k = "aaaa"	
	}()
    return
}

func test2() (string){
	k := "bbbb"
	defer func (){
		k = "cccc"
	}()
 	return k
}

