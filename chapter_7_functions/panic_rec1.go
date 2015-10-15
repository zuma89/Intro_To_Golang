package main

import "fmt"

func main() {
	panic("Panic")
	str := recover()
	fmt.Println(str)
}