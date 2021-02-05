package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

//Error Handling Code
func errorHandling(msg string) (int, error) {

	return 0, errors.New("The Error has been occured for the same")

}

//Panic Handling code

func panicHandle() {
	if a := recover(); a != nil {
		fmt.Println("Panic successfully recovered")
		//Always used in defer call in same go routine at the start of function
	}
}

//Wait Groups defined
var wg sync.WaitGroup

//Single Producer
func producer(job chan<- string) {
	for _, i := range message {
		job <- i

	}

	defer wg.Done()
	defer close(job)

}

var consumerCount = 2

//Multple consumer soa s to be returned
//The consumer function is written as:

func consumer(name int, job <-chan string) {
	for i := range job {
		fmt.Printf(" name:%v has consumed message %s\n", name, i)
	}
	wg.Done()

}

//Calling of Producer and consumer
func producerConsumer() {
	var channel = make(chan string)
	wg.Add(consumerCount + 1)
	go producer(channel)

	for i := 1; i <= consumerCount; i++ {
		go consumer(i, channel)
	}
	wg.Wait()

}

//Function to control the execution
func routine(command <-chan string, masterWait *sync.WaitGroup) {
	defer masterWait.Done()
	var status = "Pause"
	for {
		select {
		case cmd := <-command:
			fmt.Println(cmd)
			switch cmd {
			case "Stop":
				return
			case "Pause":
				fmt.Printf("for 2 sec\n")
				status = "Pause"
			default:

				status = "Play"

			}
		default:
			if status == "Play" {

				producerConsumer()
			}
		}
	}

}

//Main Function

func main() {
	var masterWait sync.WaitGroup
	masterWait.Add(1)
	command := make(chan string)
	go routine(command, &masterWait)

	for k := 0; k < 1; k++ {

		command <- "Play"

		time.Sleep(1 * time.Second)
		command <- "Pause"

		time.Sleep(2 * time.Second)
	}
	time.Sleep(3 * time.Second)
	command <- "Stop"
	defer close(command)
	masterWait.Wait()

}

//Sample string for testing purpose
var message = []string{
	"The world itself's",
	"just one big hoax.",
	"Spamming each other with our",
	// 	"running commentary of bullshit,",
	// 	"masquerading as insight, our social media",
	// 	"faking as intimacy.",
	// 	"Or is it that we voted for this?",
	// 	"Not with our rigged elections,",
	// 	"but with our things, our property, our money.",
	// 	"I'm not saying anything new.",
	// 	"We all know why we do this,",
	// 	"not because Hunger Games",
	// 	"books make us happy,",
	// 	"but because we wanna be sedated.",
	// 	"Because it's painful not to pretend,",
	// 	"because we're cowards.",
	// 	"- Elliot Alderson",
	// 	"Mr. Robot",
}
