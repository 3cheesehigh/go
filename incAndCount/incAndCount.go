package main

import (
	"fmt"
	"math/rand"
	. "time"
)

var (
	Counter int
	done    chan bool
)

func v() { Sleep(Duration(rand.Int63n(1e7))) }

func inc(n *int) {
	Accu := *n //"LDA n"
	v()
	Accu++ //"INA"
	v()
	*n = Accu //"STA n"
	fmt.Println(*n)
	v()
}

func count(p int) {
	const N = 100
	for n := 0; n < N; n++ {
		inc(&Counter)
	}
	done <- true
}

func main() {
	Counter = 0
	done = make(chan bool)
	go count(0)
	go count(1)
	<-done
	<-done
	fmt.Printf("Zähler =   %d\n", Counter)
}
