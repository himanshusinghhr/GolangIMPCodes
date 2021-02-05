package main

import (
	"fmt"
	"sync"
	"time"
)

var count = 0

//Producer is Code

func routine() {
	for {
		select {
		case <-pause:
			// fmt.Println("Producer pause")
			select {
			case <-play:
				// fmt.Println("play")
			case <-quit:
				wg.Done()
				return
			}
		case <-quit:
			wg.Done()
			return
		default:
			time.Sleep(200 * time.Millisecond)
			i++

			channel <- i
		}
	}
}

var channel = make(chan int, 5)

func prod() {
	go routine()
	for {
		time.Sleep(2 * time.Second)
		pause <- struct{}{}

		time.Sleep(5 * time.Second)
		play <- struct{}{}

	}

}

//Main Function to execute

func main() {
	wg.Add(2)
	go prod()
	go consume()
	wg.Wait()
	fmt.Println("done")
}

//Consumer Code

func consume() {
	go routine2()
	for {
		time.Sleep(2 * time.Second)
		pause <- struct{}{}

		time.Sleep(5 * time.Second)
		play <- struct{}{}

	}
}
func routine2() {
	for {
		select {
		case <-pause:
			// fmt.Println("Consumer pause")
			select {
			case <-play:
				// fmt.Println("play")
			case <-quit:
				wg.Done()
				return
			}
		case <-quit:
			wg.Done()
			return
		default:
			time.Sleep(200 * time.Millisecond)

			msg := <-channel
			fmt.Printf("This message Consumed : %d\n", msg)
		}
	}
}

var k = 0
var play = make(chan struct{})
var pause = make(chan struct{})
var quit = make(chan struct{})
var wg sync.WaitGroup
var i = 0

//
