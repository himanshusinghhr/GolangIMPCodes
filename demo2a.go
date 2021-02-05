package main

import (
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"
)

var count = 0

func routine() {
	for {
		select {
		case <-pause:
			fmt.Println("pause")
			select {
			case <-play:
				fmt.Println("play")
			case <-quit:
				wg.Done()
				return
			}
		case <-quit:
			wg.Done()
			return
		default:
			work()
			// count++
			// fmt.Printf(" count :%d", count)
			// fmt.Println(<-channel)

			// go consumer(1)
			log.Println(runtime.NumGoroutine())

		}
	}
}

var consumerCount = 10

func main() {
	wg.Add(1)
	go routine()
	// log.Println(runtime.NumGoroutine())
	time.Sleep(1 * time.Second)
	pause <- struct{}{}
	// log.Println(runtime.NumGoroutine())
	time.Sleep(5 * time.Second)
	play <- struct{}{}
	// log.Println(runtime.NumGoroutine())
	time.Sleep(1 * time.Second)
	pause <- struct{}{}
	// log.Println(runtime.NumGoroutine())

	time.Sleep(1 * time.Second)
	play <- struct{}{}

	time.Sleep(1 * time.Second)
	close(quit)

	wg.Wait()
	fmt.Println("done")
}

var channel = make(chan int)

func work() {
	time.Sleep(2 * time.Second)
	count++
	fmt.Printf(" count :%d", count)
	i++
	// fmt.Println(i)
	channel <- i
	fmt.Println(<-channel)
	// i++
	// fmt.Println(i)
}

// func consumer(name int) {
// 	time.Sleep(250 * time.Millisecond)
// 	for i = range channel {
// 		// fmt.Printf("%v has consumed %d\n", name, i)
// 		i = i + 1
// 		i = i - 1
// 		// time.Sleep(time.Millisecond * 200)
// 	}

// }

var play = make(chan struct{})
var pause = make(chan struct{})
var quit = make(chan struct{})
var wg sync.WaitGroup
var i = 0
