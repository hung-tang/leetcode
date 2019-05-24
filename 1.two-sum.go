/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		ans := target - num
		idx, ok := m[ans]
		if ok {
			return []int{idx, i}
		}
		m[num] = i
	}
	return nil
}