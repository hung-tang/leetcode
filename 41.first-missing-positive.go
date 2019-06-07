package main

/*
 * @lc app=leetcode id=41 lang=golang
 *
 * [41] First Missing Positive
 */
func firstMissingPositive(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 1
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] <= 0 || nums[i] > len(nums) {
			continue
		}
		for j := i; nums[i] >= 1 && nums[i] <= len(nums) && nums[nums[j]-1] != nums[j]; {
			tmp := nums[nums[j]-1]
			nums[nums[j]-1] = nums[j]
			nums[j] = tmp
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}

/* func main() {
	fmt.Println(firstMissingPositive([]int{1}))
	fmt.Println(firstMissingPositive([]int{1, 2, 0}))
	fmt.Println(firstMissingPositive([]int{3, 4, -1, 1}))
	fmt.Println(firstMissingPositive([]int{7, 8, 9, 11, 12}))
} */
