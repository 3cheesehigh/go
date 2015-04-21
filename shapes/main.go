package main

import "fmt"
import (
	. "github.com/3cheesehigh/shapes/circle"
	. "github.com/3cheesehigh/shapes/rectangle"
	"github.com/3cheesehigh/shapes/shape"
)

func main() {
	c := Circle{10}
	r := Rectangle{3, 5}
	fmt.Println(Shape.TotalArea(&c, &r))
}
