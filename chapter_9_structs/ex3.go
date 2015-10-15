package main

import(
	"fmt"
	"math"
)

type Rectangle struct{
	x1, y1, x2, y2 float64
}

type Circle struct{
	x, y, r float64
}

type Shape interface{
	area() float64
	perimeter() float64
}

func (c *Circle) area() float64 {
    return math.Pi * c.r*c.r
}

func distance(x1, y1, x2, y2 float64) float64 { //distance from corner to corner
    a := x2 - x1
    b := y2 - y1
    return math.Sqrt(a*a + b*b)
}

func (r *Rectangle) area() float64 {
    l := distance(r.x1, r.y1, r.x1, r.y2)
    w := distance(r.x1, r.y1, r.x2, r.y1)
    return l * w
}

func (c *Circle) perimeter() float64 {
    return 2 * math.Pi * c.r
}

func (r *Rectangle) perimeter() float64 {
    l := distance(r.x1, r.y1, r.x1, r.y2)
    w := distance(r.x1, r.y1, r.x2, r.y1)
    return 2 * (l + w)
}

func main(){
	c := Circle{1, 1, 5}
	r := Rectangle{1, 1, 10, 10}
	
	fmt.Println(c.area())
	fmt.Println(r.area())
	
	fmt.Println(c.perimeter())
	fmt.Println(r.perimeter())
	
	fmt.Println(distance(1, 1, 10, 10))
}
