package main

/*
 * @lc app=leetcode id=62 lang=golang
 *
 * [62] Unique Paths
 */
func uniquePaths(m int, n int) int {
	grid := make([][]int, n)
	for r := range grid {
		grid[r] = make([]int, m)
	}
	for c := range grid[0] {
		grid[0][c] = 1
	}
	for r := range grid {
		grid[r][0] = 1
	}
	for r := 1; r < n; r++ {
		for c := 1; c < m; c++ {
			grid[r][c] = grid[r-1][c] + grid[r][c-1]
		}
	}
	return grid[n-1][m-1]
}

// func main() {
// 	fmt.Println(uniquePaths(2, 1))
// }
