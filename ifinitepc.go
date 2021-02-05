package main

import (
	"fmt"
	"sync"
	"time"
)

//Producer producing for infinite no of times
func producer(job chan<- int) {
	var abc int
	for {
		abc = abc + 1
		job <- abc
		time.Sleep(time.Second * 1)

	}
	defer close(job)
	defer wg.Done()

}

var wg sync.WaitGroup

func main() {

	channel := make(chan int)

	wg.Add(consumerCount + 1)
	go producer(channel)
	for i := 0; i < consumerCount; i++ {

		go consumer(i, channel)

	}
	wg.Wait()

}

var consumerCount = 10

//Consumer consumimg for infinite no of times
func consumer(name int, job <-chan int) {
	for i := range job {
		fmt.Printf("%v has consumed %d\n", name, i)
		time.Sleep(time.Millisecond * 200)
	}
	wg.Done()
}
