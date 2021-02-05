package main

import (
	"fmt"
)

func handlepanic() {
	if a := recover(); a != nil {

		fmt.Println("Recovered")
	}
}

func entry(lan, nam *string) {

	defer handlepanic()
	if *lan == "hindi" && *nam == "Himanshu" {
		panic("Language is hindi")

	}

}
func main() {
	var lang = "hindi"
	var name = "Himanshu"
	entry(&lang, &name)
	fmt.Println("Ended")

}
