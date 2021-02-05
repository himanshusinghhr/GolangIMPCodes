package main

import (
	"fmt"
)

func myfunc(ch chan int) {

	fmt.Println(234 + <-ch)
}
func main() {
	fmt.Println("start Main method")
	// Creating a channel
	ch := make(chan int, 2)

	go myfunc(ch)
	ch <- 23

	ch <- 24
	go myfunc(ch)
	ch <- 25
	go myfunc(ch)
	// fmt.Scanln()

	fmt.Println("End Main method")
}
