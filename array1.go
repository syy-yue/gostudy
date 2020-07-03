package main

import "fmt"

func main() {
	a := [5]int{1,2}
	c := f1(a)
	fmt.Println(c)

}

func f1(a [5]int) (b int){
	fmt.Println(a)
	b = 23
	return
}
