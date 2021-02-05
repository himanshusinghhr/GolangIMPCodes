// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	r1 := make(chan string)
// 	r2 := make(chan string)
// 	go func() {
// 		for {
// 			r1 <- "I am running every 500 ms"
// 			time.Sleep(time.Millisecond * 500)
// 		}
// 	}()
// 	go func() {
// 		for {
// 			r2 <- "I am running every 2 sec"
// 			time.Sleep(time.Second * 2)
// 		}
// 	}()
// 	for {
// 		select {
// 		case mesg1 := <-r1:
// 			fmt.Println(mesg1)
// 		case mesg2 := <-r2:
// 			fmt.Println(mesg2)
// 		default:
// 			fmt.Println("Nothing is found")

// 		}
// 		// fmt.Println(<-r1)
// 		// fmt.Println(<-r2)
// 	}

// }
