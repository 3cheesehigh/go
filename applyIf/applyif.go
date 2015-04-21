package main 



type function func(i int ) int 
type predicate func (i int ) bool 

func applyIf(p predicate, f function, xs []int) []int{
	ys := make([]int, 0)
	for _, x := range xs {
		ys = append(ys, f(x))
	}
	return ys
}

func even(i int) bool {
	return i % 2 == 0 
}

func square(i int) int {
	return i * i
}

func main () {
	for _, x := range applyIf(even, square, []int{1,2,3,4,5,6,7}){
		println(x)
	}
}
