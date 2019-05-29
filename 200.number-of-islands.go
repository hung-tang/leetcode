package main

/*
 * @lc app=leetcode id=200 lang=golang
 *
 * [200] Number of Islands
 */
func numIslands(grid [][]byte) int {
	res := 0
	for r := range grid {
		for c := range grid[r] {
			if string(grid[r][c]) == "1" {
				res++
				dfs3(grid, r, c)
			}
		}
	}
	return res
}

func dfs3(grid [][]byte, r int, c int) {
	grid[r][c] = 0
	numRows := len(grid)
	numCols := len(grid[0])
	//top
	if r-1 >= 0 && string(grid[r-1][c]) == "1" {
		dfs3(grid, r-1, c)
	}
	//right
	if c+1 < numCols && string(grid[r][c+1]) == "1" {
		dfs3(grid, r, c+1)
	}
	//bottom
	if r+1 < numRows && string(grid[r+1][c]) == "1" {
		dfs3(grid, r+1, c)
	}
	//left
	if c-1 >= 0 && string(grid[r][c-1]) == "1" {
		dfs3(grid, r, c-1)
	}
}

/* func main() {
	var grid = [][]byte{
		{1, 1, 1, 1, 0},
		{1, 1, 0, 1, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}
	fmt.Println(numIslands(grid))
} */
