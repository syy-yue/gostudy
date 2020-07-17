package main

import(
	"fmt"
	"log"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)

func main(){
	var isOpen bool
	db,err := sql.Open("mysql","songyueyue:syy@tcp(127.0.0.1:3306)/hello")
	
	if err != nil{
		log.Fatal(err)
	}else{
		isOpen = true
	}
	
	fmt.Println(isOpen)
	db.Close()
}
