package main

import "fmt"

func half(x int) string{
	total := 0
	total = x / 2
	if total % 2 == 0 {
		return "false"
	}else {
		return "true"
	}
}

func main(){
	x := 5
	fmt.Println(half(x))
}