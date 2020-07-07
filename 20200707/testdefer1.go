package main

import "fmt"

func main(){
	for i :=0;i<3;i++{
		defer fmt.Println(i)
    }
	
	defer fmt.Println()

    for i := 0;i<3;i++{
    	defer func(){
			fmt.Println(i)
        }()
    }
}

