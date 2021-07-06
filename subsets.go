package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := []int{1, 2, 3}
	subsets(&a, 0, len(a), osf)

}

var osf = ""

func subsets(a *[]int, i, n int, osf string) {
	if i == n {
		fmt.Println(osf)
		return
	}
	subsets(a, i+1, n, osf+strconv.Itoa((*a)[i]))
	subsets(a, i+1, n, osf)

}
