package main

import "fmt"

type Person struct{
	Name string
}

func (p *Person) Talk() string{
	return ("Hi! My name is ", p.Name)
}

func main(){
	p := Person{"Hilary"}
	fmt.Println(Talk())
}