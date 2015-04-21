package Shape

type Shape interface {
	Area() float64
}

func TotalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.Area()
	}
	return area
}
