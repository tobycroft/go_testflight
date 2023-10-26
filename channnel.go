package main

import (
	"fmt"
)

var ChanFlight = make(chan any, 100)

//func main() {
//	go readout()
//	writein()
//	time.Sleep(100 * time.Millisecond)
//
//}

func readout() {
	for c := range ChanFlight {
		fmt.Println(c)
		//time.Sleep(1 * time.Millisecond)
	}
}

func writein() {
	for i := 0; i < 100; i++ {
		ChanFlight <- i
	}
}
