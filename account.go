package main

import "fmt"

func main() {	
             i := 1
            S: for ;i < 10 ;{
              j := 1
			  
			 R:           
			   
                         
			             // var sum int
                          sum :=  i * j 
						  
                        fmt.Printf("%d * %d = %d ",j,i,sum)
                          
                        if i == j{
						fmt.Printf("\n")
						i++
						goto S
}
                        j++
                        goto R

                  
                      
				      
               }
             

}

