package main

import ("fmt"
		"io/ioutil"
		"log" 
)

func main() {
	/*
	遍历文件夹
	*/
	dirname := "/home/songyueyue"
	listFiles(dirname,0)
}

func listFiles(dirname string,level int){
	//level1用来记录当前递归的层次，生成带有层次感的空格
    s := "|--"
    for i:=0;i<level;i++{
		s = "| "+s
    }
	fileInfo,err := ioutil.ReadDir(dirname)
    if err != nil{
		log.Fatal(err)
    }
	for _,fi := range fileInfo{
		fileName := dirname+"/"+fi.Name()
		fmt.Printf("%s%s\n",s,fileName)
        if fi.IsDir(){
			//递归调用方法
            listFiles(fileName,level+1)

}
}
}
