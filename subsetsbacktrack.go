package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3}
	subsets(&a, 0, len(a))

}

var osf = []int{}

func subsets(a *[]int, i, n int) {
	if i == n {
		fmt.Println(osf)
		return
	}
	osf = append(osf, (*a)[i])
	subsets(a, i+1, n)
	osf = osf[:len(osf)]
	subsets(a, i+1, n)

}
