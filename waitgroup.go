package main

import (
	"fmt"
	"sync"
)

func waitGroupImplement(wg *sync.WaitGroup) {
	fmt.Println("Hello i in wait group 1")
	wg.Done()

}
func waitGroupImplement2(wg *sync.WaitGroup) {
	fmt.Println("Hello i in wait group 2")
	wg.Done()

}
func main() {
	fmt.Println("Hello using starting main ")
	var wg sync.WaitGroup
	wg.Add(2)
	go waitGroupImplement(&wg)
	go waitGroupImplement2(&wg)
	wg.Wait()
	fmt.Println("Hello using ending main ")

}
