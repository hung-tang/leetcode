package main

/*
 * @lc app=leetcode id=63 lang=golang
 *
 * [63] Unique Paths II
 */
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 {
		return 0
	}
	//edge case 1 [[1 0] [0 0]]
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	//edge case 2 [[0]]
	if len(obstacleGrid) == 1 && len(obstacleGrid[0]) == 1 && obstacleGrid[0][0] == 0 {
		return 1
	}

	maxRows := len(obstacleGrid)
	maxCols := len(obstacleGrid[0])
	// I won't be checking obstacleGrid[0][0]
	for c := 1; c < maxCols; c++ { // set the first row
		if obstacleGrid[0][c] == 0 {
			obstacleGrid[0][c] = 1
			continue
		}
		if obstacleGrid[0][c] == 1 {
			obstacleGrid[0][c] = 0
			for i := c + 1; i < maxCols; i++ {
				obstacleGrid[0][i] = 0
			}
			break
		}
	}
	for r := 1; r < maxRows; r++ { // set the first column
		if obstacleGrid[r][0] == 0 {
			obstacleGrid[r][0] = 1
			continue
		}
		if obstacleGrid[r][0] == 1 {
			obstacleGrid[r][0] = 0
			for i := r + 1; i < maxRows; i++ {
				obstacleGrid[i][0] = 0
			}
			break
		}
	}
	for r := 1; r < maxRows; r++ {
		for c := 1; c < maxCols; c++ {
			if obstacleGrid[r][c] == 1 {
				obstacleGrid[r][c] = 0
				continue
			}
			obstacleGrid[r][c] = obstacleGrid[r-1][c] + obstacleGrid[r][c-1]
		}
	}
	return obstacleGrid[maxRows-1][maxCols-1]
}

// func main() {
// 	grid := [][]int{
// 		{0, 0, 0},
// 		{0, 1, 0},
// 		{0, 0, 0},
// 	}
// 	fmt.Println(uniquePathsWithObstacles(grid))
// }
