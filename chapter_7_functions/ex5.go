package main

import "fmt"

func fib(n int) int{
	if n == 0 {
		return 0
	}else if n == 1 {
		return 1
	}else {
		return (n - 1) + (n - 2)
	}
}

func main(){
	fmt.Println(fib(2))
	fmt.Println(fib(5))
	fmt.Println(fib(6))
	fmt.Println(fib(10))
}