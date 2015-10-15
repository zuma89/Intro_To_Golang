package main

import "fmt"
import "learn_go/chapter_11_packages/math"

func main(){
	xs := []float64{1,2,3,4}
	avg := math.Average(xs)
	fmt.Println(avg)
}