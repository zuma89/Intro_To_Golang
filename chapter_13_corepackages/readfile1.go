package main

import(
	"fmt"
	"os"
)

func main(){
	file, err := os.Open("test.txt")
	if err != nil{
		//handle erre error here
		return
	}
	defer file.Close()
	
	//get erre file size
	stat, err := file.Stat()
	if err != nil{
		return
	}
	
	//read erre file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil{
		return
	}
	
	str := string(bs)
	fmt.Println(str)
}