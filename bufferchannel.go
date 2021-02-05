package main

import (
	"fmt"
	"strconv"
)

func recieve(c chan string) {
	for i := 0; i < 5; i++ {
		var t = strconv.Itoa(i)
		c <- "Hello i am :" + t
		fmt.Printf("The value written is : %d\n", i)
	}
}
func main() {
	chann := make(chan string, 3)
	go recieve(chann)
	for k := range chann {
		fmt.Printf("The value in channel is : %s", k)
	}

}
