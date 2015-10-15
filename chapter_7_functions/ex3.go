package main

import "fmt"

func max_num(args []int) int{
	max := args[0]
	for _, value := range args {
		if value > max {
			max = value
		}
	}
	return max
}
func main() {
	args := []int{
	48, 96, 86, 68,    
	57, 82, 63, 70,    
	37, 34, 83, 27,    
	19, 97, 9, 17}
	
    fmt.Println(max_num(args))
}