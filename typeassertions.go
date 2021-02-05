package main

import (
	"fmt"
)

func main() {
	var inter1 interface{} = "Hi i am himanshu"
	var inter2 interface{} = "Hi i am singh"

	var value1 = inter1.(string)
	var value2 = inter2.(string)
	fmt.Println()
	if value1 == value2 {
		fmt.Println("String matched")

	} else {
		fmt.Println("String unmatched")
	}
	// if e {
	// 	fmt.Println("Integer is there")
	// 	fmt.Println(value)
	// } else {
	// 	fmt.Println("Integer is not there")
	// }

}
