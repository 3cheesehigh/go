package main

import "fmt"

// Rekurisv
func fib(i int) int {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}

	return fib(i-1) + fib(i-2)

}

// Endrekursiv
func qfib(i, f0, f1 int) int {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return f1 + f0
	}
	return qfib(i-1, f1+f0, f0)
}

// Iterativ
func iterfib(i int) int {
	f0, f1 := 0, 1
	if i == 0 {
		return 0
	}
	for i != 1 {
		tmp := f0
		f0 = f1 + f0
		f1 = tmp
		i--
	}

	return f1 + f0
}

func main() {
	for x := 0; x < 100; x++ {
		fmt.Printf("Die %d Fibonaccizahl mod 3 lautet: %d\n", x, iterfib(x)%3)
	}
}
