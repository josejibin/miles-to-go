package main

import (
	"fmt"
	"os"
	"strconv"
)

type Circle struct {
	radius float64
}

type Rectangle struct {
	width  float64
	lenght float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius
}

func (r Rectangle) Area() float64 {
	return r.lenght * r.width
}

func main() {
	switch shape := os.Args[1]; shape {
	case "circle":
		{
			fmt.Println("Enter radius of circle")
			r := os.Args[1]
			rad, _ := strconv.ParseFloat(r, 64)
			circle := Circle{rad}
			fmt.Println(circle)
			area := circle.Area()
			fmt.Println(area)
		}
	case "rectangle":
		{
			fmt.Println("Enter lenght and width of rectangle")
			l, w := os.Args[1], os.Args[2]
			le, _ := strconv.ParseFloat(l, 64)
			wi, _ := strconv.ParseFloat(w, 64)
			rectangle := Rectangle{le, wi}
			fmt.Println(rectangle)
			area_r := rectangle.Area()
			fmt.Println(area_r)
		}
	default:
		fmt.Println(shape)
	}
}
