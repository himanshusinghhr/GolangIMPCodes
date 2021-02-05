package main

import (
	"fmt"
)

func main() {

	score := make(map[string]int)
	defer fmt.Println(score)
	score["Himanshu"] = 34
	score["Ronny1"] = 36
	score["Ronny2"] = 40
	score["Ronny3"] = 56
	delete(score, "Ronny2")
	for k, v := range score {
		fmt.Printf("The value is : %s and the value is : %d\n", k, v)
	}
	fmt.Println(score["Himanshu"])
}
