package main

/*
 * @lc app=leetcode id=239 lang=golang
 *
 * [239] Sliding Window Maximum
 */
func maxSlidingWindow(nums []int, k int) []int {
	res := make([]int, 0)
	if nums == nil || len(nums) == 0 || k == 0 {
		return res
	}
	q := make([]int, 0)
	m := make([]int, 0)
	for i := 0; i < k; i++ {
		insert239(&q, &m, k, nums[i])
	}
	res = append(res, m[0])
	for i := k; i < len(nums); i++ {
		insert239(&q, &m, k, nums[i])
		res = append(res, m[0])
	}
	return res
}

func insert239(q *[]int, m *[]int, k int, num int) {
	if len(*q) == k {
		if (*m)[0] == (*q)[0] {
			*m = (*m)[1:] // remove max
		}
		*q = (*q)[1:]
	}
	*q = append(*q, num)
	for len(*m) > 0 && (*m)[len(*m)-1] < num {
		*m = (*m)[:len(*m)-1]
	}
	*m = append(*m, num)
}

// func main() {
// 	var nums = []int{1, 3, -1, -3, 5, 3, 6, 7}
// 	k := 3
// 	fmt.Println(maxSlidingWindow(nums, k))
// }
