package main

import (
	"fmt"
	"sync"
)

var x int64
var wg sync.WaitGroup
var ch = make(chan int)

func odd() {
	for i := 0; i < 50; i++ {
		ch <- 1
		x = x + 1
		fmt.Printf("%d\n", x)
	}
	wg.Done()
}
func even() {
	for i := 0; i < 50; i++ {
		<-ch
		x = x + 1
		fmt.Printf("%d\n", x)
	}
	wg.Done()
}
func main() {

	wg.Add(2)
	go odd()
	go even()
	wg.Wait()

}
