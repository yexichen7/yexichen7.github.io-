package main

import (
	"fmt"
	"sync"
)

var x int64
var wg sync.WaitGroup
var ch = make(chan int)

func add() {
	for i := 0; i < 50000; i++ {
		ch <- 1
		x = x + 1

	}
	wg.Done()
}
func add2() {
	for i := 0; i < 50000; i++ {
		<-ch
		x = x + 1

	}
	wg.Done()

}
func main() {

	wg.Add(2)
	go add()
	go add2()
	wg.Wait()
	fmt.Printf("%d", x)
}
