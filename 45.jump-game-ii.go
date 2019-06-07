package main

/*
 * @lc app=leetcode id=45 lang=golang
 *
 * [45] Jump Game II
 */
func jump(nums []int) int {
	if nums == nil || len(nums) <= 1 { // assume can always reach the last index
		return 0
	}
	// use bfs
	pos := make([]int, len(nums))
	for i, start := 0, 1; start < len(pos); i++ {
		if nums[i] == 0 {
			continue
		}
		end := min4(i+nums[i], len(nums)-1)
		if end < start {
			continue
		}
		for j := start; j <= end; j++ {
			pos[j] = pos[i] + 1
		}
		start = end + 1
	}
	return pos[len(pos)-1]
}

func min4(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

/* func main() {
	fmt.Println(jump([]int{3, 1, 1, 1, 1}))
	fmt.Println(jump([]int{2, 0, 2, 0, 1}))
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
} */
