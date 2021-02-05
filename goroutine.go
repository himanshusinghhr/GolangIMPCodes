package main

import (
	"fmt"
	"sync"
	"time"
)

func display(s string, w *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}

	w.Done()

}
func main() {
	// go display("I am routine")

	var wg sync.WaitGroup
	wg.Add(1)
	go display("Hello", &wg)
	wg.Wait()

	fmt.Println("Hi i am ending")

	// time.Sleep(2 * time.Second)
	// fmt.Scanln()

}

// 2 ways to create a channel
// var channel_name chan types ///NIL DEfault
// x:= make(chan type)

// := var

// // := only inside fucntion or methods hogi
// var i =[]{dkfnsgi,adhhd.djfbfhd}//slice//append/delete
// var i =[3]{dkfnsgi,adhhd.djfbfhd} //this is array
// var i = make([]string)
