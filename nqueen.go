package main

import "fmt"

type index struct {
	row int
	col int
}

func main() {
	var n = 16
	grid := make(map[index]bool, n*n)
	fmt.Println(len(grid))
	countnqueen(&grid, 0, n)
	fmt.Println(count)
}

func isSafeCell(grid *map[index]bool, current_row, current_col, n int) bool {
	//col check
	for i := current_row - 1; i >= 0; i-- {
		if (*grid)[index{i, current_col}] {
			return false
		}

	}
	//left upper diaognal
	for i, j := current_row-1, current_col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if (*grid)[index{i, j}] {
			return false
		}

	}
	//right upper diagonal
	for i, j := current_row-1, current_col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if (*grid)[index{i, j}] {
			return false
		}
	}
	return true

}

var count = 0

func countnqueen(grid *map[index]bool, current_row int, n int) {
	if current_row == n {
		count++
	}

	for j := 0; j < n; j++ {
		if isSafeCell(grid, current_row, j, n) {

			(*grid)[index{current_row, j}] = true
			countnqueen(grid, current_row+1, n)
			(*grid)[index{current_row, j}] = false
		}
	}
}
