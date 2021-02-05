package main

import (
	"fmt"
)

func testChannel(c chan int) {
	fmt.Println("Hello in chan")

	c <- 257

}
func main() {
	fmt.Println("Hello i am communication channel")
	chann := make(chan int)

	// chann <- 20
	// var wg sync.WaitGroup

	go testChannel(chann)
	defer close(chann)
	message := <-chann
	fmt.Println(message)

}

//Channel me jb bhi kch write ya read krega to wo khudko wait state me daal dega
