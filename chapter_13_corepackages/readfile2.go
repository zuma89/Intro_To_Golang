package main 

import(
	"fmt"
	"io/ioutil"
)

func main(){
	bs, err := ioutil.Readfile("test.txt")
	if err != nil{
		return
	}
	str := string(bs)
	fmt.Println(str)
}