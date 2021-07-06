package main

import "fmt"

func main() {
	n := 7
	a := [][]int{
		{0, 0, 1, 0, 0, 1, 0},
		{1, 0, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 1},
		{1, 0, 1, 0, 0, 0, 0},
		{1, 0, 1, 1, 0, 1, 0},
		{1, 0, 0, 0, 0, 1, 0},
		{1, 1, 1, 1, 0, 0, 0},
	}
	visited := make(map[index]bool, n)

	helper(0, 0, 7, &visited, &a)
	fmt.Println(totalPaths)
}

type index struct {
	row int
	col int
}

func isSafe(i int, j int, n int, visited *map[index]bool, grid *[][]int) bool {
	if i >= 0 && j >= 0 && i < n && j < n && (*visited)[index{i, j}] == false {
		return true

	}
	return false

}

var totalPaths = 0

func helper(i, j, n int, visited *map[index]bool, grid *[][]int) {
	if i == n-1 && j == n-1 {
		totalPaths++
		return
	}
	//pruning
	if !isSafe(i, j, n, visited, grid) {
		return
	}
	//jis bhi cell pe ho true kro
	(*visited)[index{i, j}] = true
	if i-1 >= 0 && (*grid)[i-1][j] == 0 {
		helper(i-1, j, n, visited, grid)
	}
	if i+1 < n && (*grid)[i+1][j] == 0 {
		helper(i+1, j, n, visited, grid)
	}
	if j+1 < n && (*grid)[i][j+1] == 0 {
		helper(i, j+1, n, visited, grid)
	}
	if j-1 >= 0 && (*grid)[i][j-1] == 0 {
		helper(i, j-1, n, visited, grid)
	}
	//jate time false krdo usko
	(*visited)[index{i, j}] = false
	return

}

//Worst case scenario me
//O(3^(n2-(4n-4)*2^(4n-4))

// middle cell se max 3 calls not 4 becz 4th call piche previous visited wle me hoga not counted
// edges mtlb 4n-4 se only 2 he calls hongi at max
