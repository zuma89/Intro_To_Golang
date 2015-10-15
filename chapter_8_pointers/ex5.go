package main

import "fmt"

//without using built-in function, new

func swap(x, y *int) {
	*x = 2
	*y = 1
}

func main(){
	x := 1
	y := 2
	swap(&x, &y)
	fmt.Println(x,y)
}

// alternative using built-in function, new
/*
func swap(xPtr, yPtr *int){
	*xPtr = 2
	*yPtr = 1
}

func main(){
	xPtr := new(int)
	yPtr := new(int)
	swap(xPtr, yPtr)
	fmt.Println(*xPtr, *yPtr)
}
*/