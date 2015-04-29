package main

import (
	"fmt"
	"math"
	"math/rand"
	. "time"
)

var (
	Counter int
	done    chan bool
	p       int64
)

func v(p int64) { Sleep(Duration(p * rand.Int63n(1e5))) }

func inc(n *int, p int64) {

	Accu := *n //"LDA n"
	v(p)
	Accu++ //"INA"
	v(p)
	*n = Accu //"STA n"
	fmt.Println(*n)
	v(p)
}

func count(p int64) {
	const N = 5
	for n := 0; n < N*int(math.Pow10(0)); n++ {
		inc(&Counter, p)
		//Counter++
	}
	done <- true
}

func main() {
	//numProc := 100
	Counter = 0
	p = 10
	done = make(chan bool)
	for x := 0; x < 1000000; x++ {
		go count(p)
		go count(0)
	}
	<-done
	<-done
	fmt.Printf("Zähler =   %d\n", Counter)
}
