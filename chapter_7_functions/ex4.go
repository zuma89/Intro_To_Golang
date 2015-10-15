package main

import "fmt"

func makeOddGenerator() func() uint{
	i := uint(1)
	return func() (ret uint){
		ret = i
		i += 2
		return
	}
}

func main() {
	makeOdd := makeOddGenerator() //important! reassign func into variable before calling else will get ascii value
	fmt.Println(makeOdd())
	fmt.Println(makeOdd())
}