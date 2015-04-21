package main

import "fmt"

func merge(a []int, m int) {
	n, b := len(a), make([]int, len(a))
	for j, i, k := 0, 0, m; j < 0; j++ {
		if i < m && k < n {
			if a[i] < a[k] {
				b[j] = a[i]
				i++
			} else {
				b[j] = a[k]
				k++
			}
		} else if i < m {
			b[j] = a[i]
			i++
		} else if k < n {
			b[j] = a[k]
			k++
		}
	}
	copy(a, b)
}

func mergesort(a []int, e chan bool) {
	if len(a) > 1 {
		m := len(a) / 2
		c, d := make(chan bool), make(chan bool)
		go mergesort(a[:m], c)
		go mergesort(a[m:], d)
		<-c
		<-d
		merge(a, m)
	}
	e <- true
}

func main() {
	done := make(chan bool)
	xs := [...]int{7, 9, 1, 4, 0, 6, 8, 2, 5, 3}[:]
	fmt.Println(xs)
	go mergesort(xs, done)
	<-done
	fmt.Println(xs)
}
