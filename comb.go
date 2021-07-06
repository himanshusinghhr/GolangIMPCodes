package main

import (
	"fmt"
	"strconv"
)

func main() {
	var test = []int{3, 2, 1}
	comute(test, 0, "")
}

var str = ""

func comute(test []int, pos int, str string) {
	if len(test) == 1 {
		str = str + strconv.Itoa(test[pos])
		fmt.Println(str)

	}
	for i := 1; i <= len(test); i++ {

		comute(test[i:], pos+1, str)

	}

}
