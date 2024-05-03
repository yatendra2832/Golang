package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in Golang")

	myCh := make(chan int)
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println("channel value from the First goRoutine", <-myCh)
		wg.Done()
	}(myCh, wg)
	go func(ch chan int, wg *sync.WaitGroup) {
		myCh <- 51
		close(myCh)
		wg.Done()
	}(myCh, wg)
	wg.Wait()
}
