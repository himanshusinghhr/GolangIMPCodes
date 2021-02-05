package main

import (
	"fmt"
	"time"
)

func main() {

	Chann1 := make(chan string)
	Chann2 := make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			Chann1 <- "Channel1 is reciveing every 500 "
		}
	}()
	go func() {
		for {
			time.Sleep(time.Second * 2)
			Chann2 <- "Channel2 is reciving every 2 sec"
		}

	}()

	for {
		select {
		case msg1 := <-Chann1:
			fmt.Println(msg1)

		case msg2 := <-Chann2:
			fmt.Println(msg2)
		}
		// fmt.Println(<-Chann1)
		// fmt.Println(<-Chann2)
	}

}
