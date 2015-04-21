package Rectangle

type Rectangle struct {
	A, B float64
}

func (r *Rectangle) Area() float64 {
	return r.A * r.B
}
